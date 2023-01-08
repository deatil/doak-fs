package utils

import (
    "os"
    "os/exec"
    "strings"
    "path/filepath"
    "encoding/base64"

    "golang.org/x/crypto/bcrypt"

    "github.com/deatil/doak-fs/pkg/fs"
    "github.com/deatil/doak-fs/pkg/global"
)

// 加密
func Base64Encode(str string) string {
    return base64.StdEncoding.EncodeToString([]byte(str))
}

// 解密
func Base64Decode(str string) string {
    newStr, err := base64.StdEncoding.DecodeString(str)
    if err != nil {
        return ""
    }

    return string(newStr)
}

// 对密码进行加密
func PasswordHash(password string) string {
    bytes, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
    return string(bytes)
}

// 对比明文密码和数据库的哈希值
func PasswordCheck(password, hash string) bool {
    err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
    return err == nil
}

// 程序根目录
func BasePath() string {
    var basePath string

    if path, err := os.Getwd(); err == nil {
        // 路径进行处理，兼容单元测试程序程序启动时的奇怪路径
        if len(os.Args) > 1 && strings.HasPrefix(os.Args[1], "-test") {
            basePath = strings.Replace(strings.Replace(path, `\test`, "", 1), `/test`, "", 1)
        } else {
            basePath = path
        }

        basePath, _ = filepath.Abs(basePath)
    } else {
        basePath = ""
    }

    return basePath
}

// 程序运行文件
func RunPath() string {
    // 可执行文件的绝对路径
    path, _ := exec.LookPath(os.Args[0])

    // 绝对路径
    absPath, _ := filepath.Abs(path)

    return absPath
}

// 检测路径是否正常
func CheckFilePath(path string) bool {
    // 根目录
    rootPath := global.Conf.File.Path
    rootPath, _ = fs.Filesystem.Realpath(rootPath)

    if strings.HasPrefix(path, rootPath) {
        return true
    }

    return false
}
