package controller

import (
    "github.com/labstack/echo/v4"

    "github.com/deatil/doak-fs/pkg/response"
)

/**
 * 基础类
 *
 * @create 2022-12-30
 * @author deatil
 */
type Base struct{}

// 错误返回
func (this *Base) Error(ctx echo.Context, msg string) error {
    return response.Render(ctx, "error.html", map[string]any{
        "msg": msg,
    })
}
