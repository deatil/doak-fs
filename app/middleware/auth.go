package middleware

import (
    "github.com/labstack/echo/v4"

    "github.com/deatil/doak-fs/pkg/session"
    "github.com/deatil/doak-fs/pkg/response"
)

// 权限验证
func Auth() echo.MiddlewareFunc {
    return func(next echo.HandlerFunc) echo.HandlerFunc {
        return func(ctx echo.Context) error {
            // 未登录跳转
            userid := session.Get(ctx, "userid")
            if userid == nil {
                return response.Redirect(ctx, "/auth/login")
            }

            // 存储登录账号
            ctx.Set("username", userid)

            return next(ctx)
        }
    }
}
