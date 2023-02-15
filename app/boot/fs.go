package boot

import (
    "github.com/deatil/doak-fs/pkg/fs"
    "github.com/deatil/doak-fs/pkg/global"
)

// 初始化文件管理器
func initFs() {
    fs.AddDriver("local", func() fs.IFs {
        return fs.NewLocal(global.Conf.File.Path)
    })

    driver := fs.GetDriver(global.Conf.File.Driver)

    if driver == nil {
        panic("fs driver not exists")
    }

    // 文件管理器
    global.Fs = fs.New(driver)
}
