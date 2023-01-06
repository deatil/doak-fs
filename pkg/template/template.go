package template

import (
    "io"
    "os"
    "fmt"
    "sync"
    "bytes"
    "errors"
    "strings"
    "path/filepath"

    "github.com/flosch/pongo2/v6"
    "github.com/labstack/echo/v4"
)

// NewTemplate creates a new Template struct.
func NewTemplate() *Template {
    p := &Template{
        dirs: []string{
            "view",
        },
        mutex: &sync.Mutex{},
        useEmbed: false,
    }
    p.templates = pongo2.NewSet("templates", p)

    return p
}

// ContextProcessorFunc signature.
type ContextProcessorFunc = func(echoCtx echo.Context, pongoCtx pongo2.Context)

type EmbedReadFileFunc = func(name string) ([]byte, error)

// Template implements custom pongo2 rendering engine for echo.
type Template struct {
    dirs              []string
    templates         *pongo2.TemplateSet
    contextProcessors []ContextProcessorFunc
    mutex             *sync.Mutex
    useEmbed          bool
    embedReadFileFunc EmbedReadFileFunc
}

// UseContextProcessor adds context processor to the pipeline.
func (this *Template) UseContextProcessor(processor ContextProcessorFunc) {
    this.contextProcessors = append(this.contextProcessors, processor)
}

// Abs returns absolute path to file requested.
// Search path is configured in AddDirectory method.
// And default directory is "./templates".
func (this *Template) Abs(base, name string) string {
    if filepath.IsAbs(name) {
        return name
    }

    for _, dir := range this.dirs {
        fullpath := filepath.Join(dir, name)
        _, err := os.Stat(fullpath)
        if err == nil {
            return fullpath
        }
    }

    return filepath.Join(filepath.Dir(base), name)
}

// Get reads the path's content from your local filesystem.
func (this *Template) Get(path string) (io.Reader, error) {
    var buf []byte
    var err error

    if this.useEmbed {
        if this.embedReadFileFunc == nil {
            return nil, errors.New("Embed func is not exists.")
        }

        path = strings.Replace(path, "\\", "/", -1)

        // embed 重新处理一次
        for _, dir := range this.dirs {
            fullpath := fmt.Sprintf("%s/%s", dir, path)

            buf, err = this.embedReadFileFunc(fullpath)
            if err == nil {
                return bytes.NewReader(buf), nil
            }
        }

        buf, err = this.embedReadFileFunc(path)
    } else {
        buf, err = os.ReadFile(path)
    }

    if err != nil {
        return nil, err
    }

    return bytes.NewReader(buf), nil
}

// AddDirectory adds a directory to the list of directories searched for templates.
// Default directory is "./templates".
func (this *Template) AddDirectory(dir string) {
    this.dirs = append(this.dirs, dir)
}

// 是否使用打包文件
func (this *Template) SetUseEmbed(useEmbed bool) {
    this.useEmbed = useEmbed
}

// 设置读取函数
func (this *Template) SetEmbedReadFileFunc(embedReadFileFunc EmbedReadFileFunc) {
    this.embedReadFileFunc = embedReadFileFunc
}

// 添加方法
func (this *Template) AddFuncs(fns map[string]any) {
    funcs := make(pongo2.Context)

    for name, fn := range fns {
        funcs[name] = fn
    }

    this.templates.Globals.Update(funcs)
}

// RegisterTag registers a custom tag.
// It calls pongo2.RegisterTag method.
func (this *Template) RegisterTag(name string, parserFunc pongo2.TagParser) {
    pongo2.RegisterTag(name, parserFunc)
}

// RegisterFilter registers a custom filter.
// It calls pongo2.RegisterFilter method.
func (this *Template) RegisterFilter(name string, fn pongo2.FilterFunction) {
    pongo2.RegisterFilter(name, fn)
}

// SetDebug sets debug mode to the template set.
// See pongo2.TemplateSet.Debug for more information.
func (this *Template) SetDebug(v bool) {
    this.mutex.Lock()
    defer this.mutex.Unlock()
    this.templates.Debug = v
}

// Render renders the view.
// Many other times, this is called in your echo handler functions.
func (this *Template) Render(w io.Writer, name string, data interface{}, ctx echo.Context) error {
    tmpl, err := this.templates.FromCache(name)
    if err != nil {
        return err
    }

    d, ok := data.(map[string]interface{})
    if !ok {
        return errors.New("Incorrect data format. Should be map[string]interface{}")
    }

    // run context processors
    for _, processor := range this.contextProcessors {
        processor(ctx, d)
    }

    return tmpl.ExecuteWriter(pongo2.Context(d), w)
}
