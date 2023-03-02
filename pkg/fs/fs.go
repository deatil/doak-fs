package fs

import (
    "io"
)

// 接口
type IFs interface {
    LsFile(directory string) []map[string]any
    LsDir(directory string) []map[string]any

    Read(path string) map[string]any
    Delete(paths ...string) error
    Get(path string) (string, error)
    Put(path string, contents string) error

    CreateFile(path string) error
    CreateDir(path string) error
    Upload(rd io.Reader, path string, name string) error

    Rename(oldName string, newName string) error
    Move(oldName string, newName string) error
    Copy(oldName string, newName string) error

    Exists(path string) bool
    IsFile(path string) bool
    IsDirectory(path string) bool

    FormatFile(path string) (string, error)
}

/**
 * 文件管理器
 *
 * @create 2023-2-14
 * @author deatil
 */
type Fs struct {
    Driver IFs
}

func New(driver IFs) Fs {
    return Fs{
        Driver: driver,
    }
}

// 列出文件
func (this Fs) LsFile(directory string) []map[string]any {
    files := this.Driver.LsFile(directory)
    if len(files) == 0 {
        return files
    }

    res := make([]map[string]any, 0)
    for _, file := range files {
        fileName := file["name"].(string)
        time := file["time"].(int64)

        namesmall := Filesystem.Basename(fileName)
        ext       := Filesystem.Extension(fileName)

        res = append(res, map[string]any{
            "name":      fileName,
            "namesmall": namesmall,
            "isDir":     false,
            "size":      file["size"],
            "time":      FormatTime(time),
            "type":      DetectFileType(fileName),
            "ext":       ext,
            "perm":      file["perm"],
            "permInt":   file["permInt"],
        })
    }

    return res
}

// 列出文件夹
func (this Fs) LsDir(directory string) []map[string]any {
    dirs := this.Driver.LsDir(directory)
    if len(dirs) == 0 {
        return dirs
    }

    res := make([]map[string]any, 0)
    for _, dir := range dirs {
        dirName := dir["name"].(string)
        time := dir["time"].(int64)

        namesmall := Filesystem.Basename(dirName)

        res = append(res, map[string]any{
            "name":      dirName,
            "namesmall": namesmall,
            "isDir":     true,
            "size":      dir["size"],
            "time":      FormatTime(time),
            "type":      "folder",
            "ext":       "",
            "perm":      dir["perm"],
            "permInt":   dir["permInt"],
        })
    }

    return res
}

// 读取数据
func (this Fs) Read(path string) map[string]any {
    data := this.Driver.Read(path)

    if len(data) == 0 {
        return data
    }

    dataName := data["name"].(string)

    typ   := "folder"
    ext   := ""
    isDir := true

    if this.Driver.IsFile(path) {
        typ   = DetectFileType(dataName)
        ext   = Filesystem.Extension(dataName)
        isDir = false
    }

    namesmall := Filesystem.Basename(dataName)
    time := data["time"].(int64)

    res := map[string]any{
        "name":      dataName,
        "namesmall": namesmall,
        "isDir":     isDir,
        "size":      data["size"],
        "time":      FormatTime(time),
        "type":      typ,
        "ext":       ext,
        "perm":      data["perm"],
        "permInt":   data["permInt"],
    }

    return res
}

func (this Fs) Delete(paths ...string) error {
    return this.Driver.Delete(paths...)
}

func (this Fs) Exists(path string) bool {
    return this.Driver.Exists(path)
}

func (this Fs) IsFile(path string) bool {
    return this.Driver.IsFile(path)
}

func (this Fs) IsDirectory(path string) bool {
    return this.Driver.IsDirectory(path)
}

func (this Fs) Get(path string) (string, error) {
    return this.Driver.Get(path)
}

func (this Fs) Put(path string, contents string) error {
    return this.Driver.Put(path, contents)
}

func (this Fs) CreateFile(path string) error {
    return this.Driver.CreateFile(path)
}

func (this Fs) Upload(src io.Reader, path string, name string) error {
    return this.Driver.Upload(src, path, name)
}

func (this Fs) Rename(oldName string, newName string) error {
    return this.Driver.Rename(oldName, newName)
}

func (this Fs) Move(oldName string, newName string) error {
    return this.Driver.Move(oldName, newName)
}

func (this Fs) Copy(oldName string, newName string) error {
    return this.Driver.Copy(oldName, newName)
}

func (this Fs) FormatFile(path string) (string, error) {
    return this.Driver.FormatFile(path)
}

func (this Fs) CreateDir(path string) error {
    return this.Driver.CreateDir(path)
}

// 名称
func (this Fs) Basename(path string) string {
    return Filesystem.Basename(path)
}

// 是否为文件夹
func (this Fs) ParentPath(path string) string {
    if path == "" || path == "/" {
        return ""
    }

    parentPath := Filesystem.Dirname(path)
    parentPath = Filesystem.ToSlash(parentPath)

    return parentPath
}

// 后缀
func (this Fs) Extension(path string) string {
    return Filesystem.Extension(path)
}
