package boot

import (
    "github.com/deatil/doak-fs/pkg/fs"
    "github.com/deatil/doak-fs/pkg/global"
)

// 初始化文件管理器
func initFs() {
    driverName := global.Conf.File.Driver

    var driver fs.IFs

    switch driverName {
        case "local":
            driver = fs.NewLocal()
    }

    // 文件管理器
    global.Fs = fs.New(driver)
}
