package fs

import (
    "os"
    "io"
    "fmt"
    "errors"
    "strings"
)

func NewLocal(rootPath string) Local {
    return Local{rootPath}
}

/**
 * 本地文件管理
 *
 * @create 2023-2-14
 * @author deatil
 */
type Local struct {
    rootPath string
}

// 列出文件及文件夹
func (this Local) Ls(directory string) []map[string]any {
    res := make([]map[string]any, 0)

    directory = this.formatPath(directory)

    if !this.checkFilePath(directory) {
        return res
    }

    directories, _ := Filesystem.Directories(directory)
    res = append(res, formatDirectories(directories, directory)...)

    files, _ := Filesystem.Files(directory)
    res = append(res, formatFiles(files, directory)...)

    return res
}

// 列出文件夹
func (this Local) LsDir(directory string) []map[string]any {
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
func (this Local) Read(path string) map[string]any {
    path = this.formatPath(path)

    if !this.checkFilePath(path) {
        return make(map[string]any)
    }

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
        typ   = "folder"
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
        this.deletePath(path)
    }

    return nil
}

// 删除
func (this Local) deletePath(path string) error {
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
func (this Local) Rename(oldName string, newName string) error {
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
func (this Local) Move(oldName string, newName string) error {
    oldName = this.formatPath(oldName)

    oldBasename := this.Basename(oldName)
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
func (this Local) Get(path string) (string, error) {
    path = this.formatPath(path)

    if !this.checkFilePath(path) {
        return "", errors.New("访问错误")
    }

    if !this.IsFile(path) {
        return "", errors.New("打开的不是文件")
    }

    data, err := Filesystem.Get(path)
    if err != nil {
        return "", errors.New("打开文件失败")
    }

    return data, nil
}

// 覆盖
func (this Local) Put(path string, contents string) error {
    path = this.formatPath(path)

    if !this.checkFilePath(path) {
        return errors.New("访问错误")
    }

    if !this.IsFile(path) {
        return errors.New("要更新的不是文件")
    }

    err := Filesystem.Put(path, contents)
    if err != nil {
        return errors.New("更新文件失败")
    }

    return nil
}

// 设置权限
func (this Local) CreateFile(path string) error {
    path = this.formatPath(path)

    if !this.checkFilePath(path) {
        return errors.New("访问错误")
    }

    if this.IsFile(path) {
        return errors.New("文件已经存在")
    }

    return Filesystem.Touch(path)
}

// 创建文件夹
func (this Local) CreateDir(path string) error {
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
func (this Local) Upload(src io.Reader, path string, name string) error {
    path = this.formatPath(path, name)

    if !this.checkFilePath(path) {
        return errors.New("访问错误")
    }

    if this.IsFile(path) {
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

// 是否为文件夹
func (this Local) Basename(path string) string {
    return Filesystem.Basename(path)
}

// 是否为文件夹
func (this Local) ParentPath(path string) string {
    if path == "" || path == "/" {
        return ""
    }

    parentPath := Filesystem.Dirname(path)
    parentPath = Filesystem.ToSlash(parentPath)

    return parentPath
}

func (this Local) Extension(path string) string {
    return Filesystem.Extension(path)
}

func (this Local) FormatFile(path string) (string, error) {
    path = this.formatPath(path)

    if !this.checkFilePath(path) {
        return "", errors.New("访问错误")
    }

    if !this.IsFile(path) {
        return "", errors.New("打开的不是文件")
    }

    return path, nil
}

// 检测路径是否正常
func (this Local) checkFilePath(path string) bool {
    // 根目录
    rootPath, _ := Filesystem.Realpath(this.rootPath)

    if strings.HasPrefix(path, rootPath) {
        return true
    }

    return false
}

// 格式化路径
func (this Local) formatPath(paths ...string) string {
    filePath, _ := Filesystem.Realpath(this.rootPath)

    for _, path := range paths {
        filePath = Filesystem.Join(filePath, path)
    }

    return filePath
}
