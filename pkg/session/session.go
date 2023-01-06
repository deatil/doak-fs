package session

import (
    "github.com/gorilla/sessions"
    "github.com/labstack/echo/v4"
    "github.com/labstack/echo-contrib/session"

    "github.com/deatil/doak-fs/pkg/global"
)

// 中间件
func SessionMiddleware(e *echo.Echo) {
    secret := []byte(global.Conf.Session.Secret)

    e.Use(session.Middleware(sessions.NewCookieStore(secret)))
}

// session
func Session(ctx echo.Context) *sessions.Session {
    sess, _ := session.Get(global.Conf.Session.Key, ctx)
    sess.Options = &sessions.Options{
        Path:     global.Conf.Session.Path,
        MaxAge:   global.Conf.Session.MaxAge, // 86400 * 7,
        HttpOnly: global.Conf.Session.HttpOnly,
    }

    return sess
}

// 存储
func Set(ctx echo.Context, key string, value any) error {
    sess := Session(ctx)

    sess.Values[key] = value
    return sess.Save(ctx.Request(), ctx.Response())
}

// 获取
func Get(ctx echo.Context, key string) any {
    sess := Session(ctx)

    if v, ok := sess.Values[key]; ok {
        return v
    }

    return nil
}

// 删除
func Delete(ctx echo.Context, key string) error {
    sess := Session(ctx)

    delete(sess.Values, key)

    return sess.Save(ctx.Request(), ctx.Response())
}
