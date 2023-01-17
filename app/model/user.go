package model

import (
    "github.com/deatil/doak-fs/pkg/global"
)

// 账号密码
func GetUserPassword(name string) string {
    return global.Conf.User.GetUserPassword(name)
}
