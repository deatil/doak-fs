package resources

import (
    "embed"
    "io/fs"
    "net/http"
)

//go:embed static
var Static embed.FS

//go:embed view
var View embed.FS

//go:embed config
var Config embed.FS

// 静态文件
func StaticAssets() http.FileSystem {
    // return http.FS(os.DirFS("static"))

    fsys, err := fs.Sub(Static, "static")
    if err != nil {
        panic(err)
    }

    return http.FS(fsys)
}

// 读取模板
func ReadViewFile(name string) ([]byte, error) {
    return View.ReadFile(name)
}

// 读取配置
func ReadConfig(name string) ([]byte, error) {
    return Config.ReadFile(name)
}
