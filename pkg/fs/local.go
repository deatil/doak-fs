package fs

import (
    "fmt"
)

func NewLocal() Local {
    return Local{}
}

/**
 * 本地文件管理
 *
 * @create 2023-2-14
 * @author deatil
 */
type Local struct {}

// 列出文件及文件夹
func (this Local) Ls(directory string) []map[string]any {
    res := make([]map[string]any, 0)

    directories, _ := Filesystem.Directories(directory)
    res = append(res, formatDirectories(directories, directory)...)

    files, _ := Filesystem.Files(directory)
    res = append(res, formatFiles(files, directory)...)

    return res
}

// 列出文件夹
func (this Local) LsDir(directory string) []map[string]any {
    res := make([]map[string]any, 0)

    directories, _ := Filesystem.Directories(directory)
    res = append(res, formatDirectories(directories, directory)...)

    return res
}

// 详情
func (this Local) Read(path string) map[string]any {
    size := "-"
    typ := ""
    ext := ""
    isDir := false

    if Filesystem.IsFile(path) {
        typ   = detectFileType(path)
        size  = formatSize(Filesystem.Size(path))
        ext   = Filesystem.Extension(path)
        isDir = false
    } else {
        typ   = "dir"
        isDir = true
    }

    perm, _ := Filesystem.PermString(path)
    permInt, _ := Filesystem.Perm(path)

    res := map[string]any{
        "name":      path,
        "namesmall": Filesystem.Basename(path),
        "isDir":     isDir,
        "size":      size,
        "time":      formatTime(Filesystem.LastModified(path)),
        "type":      typ,
        "ext":       ext,
        "perm":      perm,
        "permInt":   fmt.Sprintf("%o", permInt),
    }

    return res
}

// 删除
func (this Local) Delete(paths ...string) error {
    for _, path := range paths {
        if Filesystem.IsFile(path) {
            Filesystem.Delete(path)
        } else {
            Filesystem.DeleteDirectory(path)
        }
    }

    return nil
}

// 判断
func (this Local) Exists(path string) bool {
    return Filesystem.Exists(path)
}

// 是否为文件
func (this Local) IsFile(path string) bool {
    return Filesystem.IsFile(path)
}

// 是否为文件夹
func (this Local) IsDirectory(path string) bool {
    return Filesystem.IsDirectory(path)
}

// 获取
func (this Local) Get(path string, lock ...bool) (string, error) {
    return Filesystem.Get(path, lock...)
}

// 覆盖
func (this Local) Put(path string, contents string, lock ...bool) error {
    return Filesystem.Put(path, contents, lock...)
}

// 设置权限
func (this Local) Chmod(path string, mode uint32) error {
    return Filesystem.Chmod(path, mode)
}
