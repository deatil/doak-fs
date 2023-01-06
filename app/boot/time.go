package boot

import (
    "github.com/deatil/doak-fs/pkg/time"
    "github.com/deatil/doak-fs/pkg/global"
)

// 初始化时间
func initTime() {
    global.StartTime = time.Now().ToTime()
}
