package fs

import (
    "os"
    "io"
    "fmt"
    "errors"
    "strings"
)

func NewDriverLocal(rootPath string) DriverLocal {
    return DriverLocal{
        rootPath: rootPath,
    }
}

/**
 * 本地文件管理
 *
 * @create 2023-2-14
 * @author deatil
 */
type DriverLocal struct {
    // 根目录
    rootPath string
}

// 列出文件及文件夹
func (this DriverLocal) LsFile(directory string) []map[string]any {
    res := make([]map[string]any, 0)

    directory = this.formatPath(directory)

    if !this.checkFilePath(directory) {
        return res
    }

    files, _ := Filesystem.Files(directory)
    res = append(res, formatFiles(files, directory)...)

    return res
}

// 列出文件夹
func (this DriverLocal) LsDir(directory string) []map[string]any {
    res := make([]map[string]any, 0)

    directory = this.formatPath(directory)

    if !this.checkFilePath(directory) {
        return res
    }

    directories, _ := Filesystem.Directories(directory)
    res = append(res, formatDirectories(directories, directory)...)

    return res
}

// 详情
func (this DriverLocal) Read(path string) map[string]any {
    path = this.formatPath(path)

    if !this.checkFilePath(path) {
        return make(map[string]any)
    }

    size  := int64(0)

    if Filesystem.IsFile(path) {
        size = Filesystem.Size(path)
    }

    time       := Filesystem.LastModified(path)
    perm, _    := Filesystem.PermString(path)
    permInt, _ := Filesystem.Perm(path)

    res := map[string]any{
        "name":      path,
        "size":      size,
        "time":      time,
        "perm":      perm,
        "permInt":   fmt.Sprintf("%o", permInt),
    }

    return res
}

// 删除
func (this DriverLocal) Delete(paths ...string) error {
    for _, path := range paths {
        this.deletePath(path)
    }

    return nil
}

// 删除
func (this DriverLocal) deletePath(path string) error {
    path = this.formatPath(path)

    if !this.checkFilePath(path) {
        return errors.New("path error.")
    }

    if Filesystem.IsFile(path) {
        Filesystem.Delete(path)
    } else {
        Filesystem.DeleteDirectory(path)
    }

    return nil
}

// 重命名
func (this DriverLocal) Rename(oldName string, newName string) error {
    oldName = this.formatPath(oldName)
    newName = this.formatPath(newName)

    if !this.checkFilePath(oldName) {
        return errors.New("old name error.")
    }

    if !this.checkFilePath(newName) {
        return errors.New("new name error.")
    }

    if !this.Exists(oldName) {
        return errors.New("旧名称不存在")
    }

    if this.Exists(newName) {
        return errors.New("新名称已经存在")
    }

    return Filesystem.Move(oldName, newName)
}

// 移动
func (this DriverLocal) Move(oldName string, newName string) error {
    oldName = this.formatPath(oldName)

    oldBasename := Filesystem.Basename(oldName)
    newName = this.formatPath(newName, oldBasename)

    if !this.checkFilePath(oldName) {
        return errors.New("访问错误")
    }

    if !this.checkFilePath(newName) {
        return errors.New("访问错误")
    }

    if !this.Exists(oldName) {
        return errors.New("旧名称不存在")
    }

    if this.Exists(newName) {
        return errors.New("新名称已经存在")
    }

    err := Filesystem.Move(oldName, newName)
    if err != nil {
        return errors.New("移动文件失败")
    }

    return nil
}

// 复制
func (this DriverLocal) Copy(oldName string, newName string) error {
    oldName = this.formatPath(oldName)

    oldBasename := Filesystem.Basename(oldName)
    newName = this.formatPath(newName, oldBasename)

    if !this.checkFilePath(oldName) {
        return errors.New("访问错误")
    }

    if !this.checkFilePath(newName) {
        return errors.New("访问错误")
    }

    if !this.Exists(oldName) {
        return errors.New("旧名称不存在")
    }

    if this.Exists(newName) {
        return errors.New("新名称已经存在")
    }

    if this.IsDirectory(oldName) {
        err := Filesystem.CopyDirectory(oldName, newName)
        if err != nil {
            return errors.New("复制文件夹失败")
        }
    } else {
        err := Filesystem.Copy(oldName, newName)
        if err != nil {
            return errors.New("复制文件失败")
        }
    }

    return nil
}

// 获取
func (this DriverLocal) Get(path string) (string, error) {
    path = this.formatPath(path)

    if !this.checkFilePath(path) {
        return "", errors.New("访问错误")
    }

    if !Filesystem.IsFile(path) {
        return "", errors.New("打开的不是文件")
    }

    data, err := Filesystem.Get(path)
    if err != nil {
        return "", errors.New("打开文件失败")
    }

    return data, nil
}

// 覆盖
func (this DriverLocal) Put(path string, contents string) error {
    path = this.formatPath(path)

    if !this.checkFilePath(path) {
        return errors.New("访问错误")
    }

    if !Filesystem.IsFile(path) {
        return errors.New("要更新的不是文件")
    }

    err := Filesystem.Put(path, contents)
    if err != nil {
        return errors.New("更新文件失败")
    }

    return nil
}

// 设置权限
func (this DriverLocal) CreateFile(path string) error {
    path = this.formatPath(path)

    if !this.checkFilePath(path) {
        return errors.New("访问错误")
    }

    if Filesystem.IsFile(path) {
        return errors.New("文件已经存在")
    }

    return Filesystem.Touch(path)
}

// 创建文件夹
func (this DriverLocal) CreateDir(path string) error {
    path = this.formatPath(path)

    if !this.checkFilePath(path) {
        return errors.New("访问错误")
    }

    if this.IsDirectory(path) {
        return errors.New("文件夹已经存在")
    }

    return Filesystem.MakeDirectory(path, 0640, true)
}

// 上传
func (this DriverLocal) Upload(src io.Reader, path string, name string) error {
    path = this.formatPath(path, name)

    if !this.checkFilePath(path) {
        return errors.New("访问错误")
    }

    if Filesystem.IsFile(path) {
        return errors.New("文件已经存在")
    }

    // 创建文件
    dst, err := os.Create(path)
    if err != nil {
        return errors.New("创建文件没有权限")
    }
    defer dst.Close()

    // 保存
    if _, err = io.Copy(dst, src); err != nil {
        return errors.New("上传文件失败")
    }

    return nil
}

// 判断
func (this DriverLocal) Exists(path string) bool {
    path = this.formatPath(path)

    return Filesystem.Exists(path)
}

// 是否为文件
func (this DriverLocal) IsFile(path string) bool {
    path = this.formatPath(path)

    return Filesystem.IsFile(path)
}

// 是否为文件夹
func (this DriverLocal) IsDirectory(path string) bool {
    path = this.formatPath(path)

    return Filesystem.IsDirectory(path)
}

func (this DriverLocal) FormatFile(path string) (string, error) {
    path = this.formatPath(path)

    if !this.checkFilePath(path) {
        return "", errors.New("访问错误")
    }

    if !Filesystem.IsFile(path) {
        return "", errors.New("打开的不是文件")
    }

    return path, nil
}

// 检测路径是否正常
func (this DriverLocal) checkFilePath(path string) bool {
    rootPath, _ := Filesystem.Realpath(this.rootPath)

    if strings.HasPrefix(path, rootPath) {
        return true
    }

    return false
}

// 格式化路径
func (this DriverLocal) formatPath(paths ...string) string {
    filePath, _ := Filesystem.Realpath(this.rootPath)

    for _, path := range paths {
        filePath = Filesystem.Join(filePath, path)
    }

    return filePath
}

// 格式化文件
func formatFiles(files []string, path string) []map[string]any {
    res := make([]map[string]any, 0)

    for _, file := range files {
        file = Filesystem.Join(path, file)

        size      := Filesystem.Size(file)
        time      := Filesystem.LastModified(file)

        perm, _    := Filesystem.PermString(file)
        permInt, _ := Filesystem.Perm(file)

        res = append(res, map[string]any{
            "name":    file,
            "size":    size,
            "time":    time,
            "perm":    perm,
            "permInt": fmt.Sprintf("%o", permInt),
        })
    }

    return res
}

// 格式化文件夹
func formatDirectories(dirs []string, path string) []map[string]any {
    res := make([]map[string]any, 0)

    for _, dir := range dirs {
        dir = Filesystem.Join(path, dir)

        time      := Filesystem.LastModified(dir)

        perm, _    := Filesystem.PermString(dir)
        permInt, _ := Filesystem.Perm(dir)

        res = append(res, map[string]any{
            "name":    dir,
            "size":    "-",
            "time":    time,
            "perm":    perm,
            "permInt": fmt.Sprintf("%o", permInt),
        })
    }

    return res
}
