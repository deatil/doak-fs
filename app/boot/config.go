package boot

import (
    "github.com/deatil/doak-fs/pkg/fs"
    "github.com/deatil/doak-fs/pkg/global"
    "github.com/deatil/doak-fs/pkg/config"
    "github.com/deatil/doak-fs/app/resources"
)

var defaultConfFile = "./config.toml"

// 初始化配置信息
func initConfig() {
    // 默认配置
    defaultConf, _ := resources.ReadConfig("config/config.toml")

    // 配置
    if global.ConfigFile != "" {
        defaultConfFile = global.ConfigFile

        // 检测
        if !fs.Filesystem.Exists(defaultConfFile) {
            fs.Filesystem.Put(defaultConfFile, string(defaultConf))
        }

        // 读取配置
        conf, err := config.ReadConfig(defaultConfFile)
        if err != nil {
            panic(err.Error())
        }

        // 设置配置信息
        global.Conf = conf
    } else {
        conf, err := config.ReadConfigByte(defaultConf)
        if err != nil {
            panic(err.Error())
        }

        // 设置配置信息
        global.Conf = conf
    }
}
