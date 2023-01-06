package controller

import (
    "github.com/labstack/echo/v4"

    "github.com/deatil/doak-fs/pkg/time"
    "github.com/deatil/doak-fs/pkg/utils"
    "github.com/deatil/doak-fs/pkg/global"
    "github.com/deatil/doak-fs/pkg/response"
)

/**
 * 首页
 *
 * @create 2022-12-30
 * @author deatil
 */
type Index struct{
    Base
}

// 首页
func (this *Index) Index(ctx echo.Context) error {
    username := ctx.Get("username")

    startTime := time.FromTime(global.StartTime).ToDateTimeString()

    return response.Render(ctx, "index_index.html", map[string]any{
        "username": username,
        "startTime": startTime,
        "conf": global.Conf,
        "run_path": utils.RunPath(),
    })
}

