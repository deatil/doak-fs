package view

import (
    "github.com/deatil/doak-fs/app/url"

    "github.com/deatil/doak-fs/pkg/time"
    "github.com/deatil/doak-fs/pkg/utils"
    "github.com/deatil/doak-fs/pkg/config"
    "github.com/deatil/doak-fs/pkg/global"
)

// 模板方法
func ViewFuncs() map[string]any {
    funcs := make(map[string]any)

    // 配置
    funcs["getCfg"] = func() config.Conf {
        return global.Conf
    }

    // 静态文件
    funcs["assets"] = url.Assets

    // 时间
    funcs["nowTime"] = time.Now
    funcs["formatTime"] = time.FromTime
    funcs["formatTimestamp"] = time.FromTimestamp
    funcs["parseTime"] = time.Parse
    funcs["mustParseTime"] = time.MustParse

    // 图标
    funcs["faIcon"] = utils.GetFaIcon

    return funcs
}
