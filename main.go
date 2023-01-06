package main

import (
    "github.com/deatil/doak-fs/app/boot"
)

// 启动
// 打包配置文件
// go run main.go

// 使用自定义配置文件
// go run main.go --config=config.toml

// 使用模板位置 './template'
// go run main.go --view=template
func main() {
    boot.Start()
}
