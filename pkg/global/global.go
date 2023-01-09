package global

import (
    "time"

    "github.com/deatil/doak-fs/pkg/config"
)

// 配置信息
var Conf config.Conf

// 启动时间
var StartTime time.Time

// 配置文件
var ConfigFile string

// 模板位置
var ViewPath string

// 是否只使用打包文件
var IsOnlyEmbed bool
