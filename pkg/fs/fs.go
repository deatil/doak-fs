package fs

import (
    "fmt"
    "regexp"

    "github.com/deatil/lakego-filesystem/filesystem"

    fs_time "github.com/deatil/doak-fs/pkg/time"
)

// 文件管理
var Filesystem *filesystem.Filesystem

// 文件类型
var fileTypes = map[string]string{
    "image": "png|jpg|jpeg|ico|gif|bmp|svg|wbmp|avif",
    "xls"  : "xls|xlt|xla|xlsx|xltx|xlsm|xltm|xlam|xlsb",
    "word" : "doc|docx|dot|dotx|docm|dotm",
    "ppt"  : "ppt|pptx|pptm",
    "pdf"  : "pdf",
    "code" : "php|js|java|python|ruby|go|c|cpp|sql|m|h|json|html|aspx",
    "zip"  : `zip|tar\.gz|rar|rpm`,
    "text" : "txt|pac|log|md",
    "audio": "mp3|wav|flac|3pg|aa|aac|ape|au|m4a|mpc|ogg",
    "video": "mkv|rmvb|flv|mp4|avi|wmv|rm|asf|mpeg",
}

// 初始化
func init() {
    Filesystem = filesystem.New()
}

// 列出文件及文件夹
func Ls(directory string) []map[string]any {
    res := make([]map[string]any, 0)

    directories, _ := Filesystem.Directories(directory)
    res = append(res, formatDirectories(directories, directory)...)

    files, _ := Filesystem.Files(directory)
    res = append(res, formatFiles(files, directory)...)

    return res
}

// 列出文件夹
func LsDir(directory string) []map[string]any {
    res := make([]map[string]any, 0)

    directories, _ := Filesystem.Directories(directory)
    res = append(res, formatDirectories(directories, directory)...)

    return res
}

// 详情
func Read(path string) map[string]any {
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
func Delete(paths ...string) error {
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
func Exists(path string) bool {
    return Filesystem.Exists(path)
}

// 是否为文件
func IsFile(path string) bool {
    return Filesystem.IsFile(path)
}

// 获取
func Get(path string, lock ...bool) (string, error) {
    return Filesystem.Get(path, lock...)
}

// 覆盖
func Put(path string, contents string, lock ...bool) error {
    return Filesystem.Put(path, contents, lock...)
}

// 设置权限
func Chmod(path string, mode uint32) error {
    return Filesystem.Chmod(path, mode)
}

// 名称
func Basename(path string) string {
    return Filesystem.Basename(path)
}

// 合并目录
func JoinPath(elem ...string) string {
    return Filesystem.Join(elem...)
}

// 格式化文件
func formatFiles(files []string, path string) []map[string]any {
    res := make([]map[string]any, 0)

    for _, file := range files {
        file = Filesystem.Join(path, file)

        perm, _ := Filesystem.PermString(file)
        permInt, _ := Filesystem.Perm(file)

        res = append(res, map[string]any{
            "name":      file,
            "namesmall": Filesystem.Basename(file),
            "isDir":     false,
            "size":      formatSize(Filesystem.Size(file)),
            "time":      formatTime(Filesystem.LastModified(file)),
            "type":      detectFileType(file),
            "ext":       Filesystem.Extension(file),
            "perm":      perm,
            "permInt":   fmt.Sprintf("%o", permInt),
        })
    }

    return res
}

// 格式化文件夹
func formatDirectories(dirs []string, path string) []map[string]any {
    res := make([]map[string]any, 0)

    for _, dir := range dirs {
        dir = Filesystem.Join(path, dir)

        perm, _ := Filesystem.PermString(dir)
        permInt, _ := Filesystem.Perm(dir)

        res = append(res, map[string]any{
            "name":      dir,
            "namesmall": Filesystem.Basename(dir),
            "isDir":     true,
            "size":      "-",
            "time":      formatTime(Filesystem.LastModified(dir)),
            "type":      "dir",
            "ext":       "",
            "perm":      perm,
            "permInt":   fmt.Sprintf("%o", permInt),
        })
    }

    return res
}

func detectFileType(file string) string {
    extension := Filesystem.Extension(file)

    for typ, regex := range fileTypes {
        result, _ := regexp.MatchString("(?i)^" + regex + "$", extension)
        if result {
            return typ
        }
    }

    return "other";
}

// 格式化时间
func formatTime(date int64) string {
    return fs_time.FromTimestamp(date).ToDateTimeString()
}

// 格式化数据大小
func formatSize(size int64) string {
    units := []string{" B", " KB", " MB", " GB", " TB", " PB"}

    s := float64(size)

    i := 0
    for ; s >= 1024 && i < len(units) - 1; i++ {
        s /= 1024
    }

    return fmt.Sprintf("%.2f%s", s, units[i])
}
