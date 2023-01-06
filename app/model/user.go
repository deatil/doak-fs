package model

import (
    "strings"
    "github.com/deatil/doak-fs/pkg/global"
)

// 账号列表
func GetUsers() map[string]string {
    names := global.Conf.User.Names

    users := make(map[string]string)

    for _, name := range names {
        newName := strings.SplitN(name, ":", 2)
        users[newName[0]] = newName[1]
    }

    return users
}

// 账号密码
func GetUserPassword(name string) string {
    users := GetUsers()

    if password, ok := users[name]; ok {
        return password
    }

    return ""
}
