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
    "image"  : "png|jpg|jpeg|ico|gif|bmp|svg|wbmp|avif",
    "xls"    : "xls|xlt|xla|xlsx|xltx|xlsm|xltm|xlam|xlsb",
    "word"   : "doc|docx|dot|dotx|docm|dotm",
    "ppt"    : "ppt|pptx|pptm",
    "pdf"    : "pdf",
    "code"   : "php|js|java|python|ruby|rs|v|go|c|cpp|sql|m|h|json|html|aspx",
    "archive": `zip|tar\.gz|rar|rpm`,
    "text"   : "txt|pac|log|md",
    "audio"  : "mp3|wav|flac|3pg|aa|aac|ape|au|m4a|mpc|ogg",
    "video"  : "mkv|rmvb|flv|mp4|avi|wmv|rm|asf|mpeg",
    "apk"    : "apk",
    "exe"    : "exe",
    "md"     : "md",
}

func DetectFileType(file string) string {
    extension := Filesystem.Extension(file)

    for typ, regex := range fileTypes {
        result, _ := regexp.MatchString("(?i)^" + regex + "$", extension)
        if result {
            return typ
        }
    }

    return "file";
}

// 格式化时间
func FormatTime(date int64) string {
    return fs_time.FromTimestamp(date).ToDateTimeString()
}

// 格式化数据大小
func FormatSize(size int64) string {
    units := []string{" B", " KB", " MB", " GB", " TB", " PB"}

    s := float64(size)

    i := 0
    for ; s >= 1024 && i < len(units) - 1; i++ {
        s /= 1024
    }

    return fmt.Sprintf("%.2f%s", s, units[i])
}
