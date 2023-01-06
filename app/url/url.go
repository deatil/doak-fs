package url

import (
    "fmt"

    "github.com/deatil/doak-fs/pkg/global"
)

// 资源
func Assets(path string) string {
    assets := global.Conf.App.Assets

    return fmt.Sprintf("%s/%s", assets, path)
}
