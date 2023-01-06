package view

import (
    "github.com/deatil/doak-fs/app/url"
    "github.com/deatil/doak-fs/pkg/time"
)

// 模板方法
func ViewFuncs() map[string]any {
    funcs := make(map[string]any)

    // 静态文件
    funcs["assets"] = url.Assets

    // 时间
    funcs["nowTime"] = time.Now
    funcs["formatTime"] = time.FromTime
    funcs["formatTimestamp"] = time.FromTimestamp
    funcs["parseTime"] = time.Parse
    funcs["mustParseTime"] = time.MustParse

    return funcs
}
