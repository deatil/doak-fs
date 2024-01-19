package webdav

import (
    "net/http"

    "golang.org/x/net/webdav"
    "github.com/labstack/echo/v4"

    "github.com/deatil/doak-fs/pkg/utils"
    "github.com/deatil/doak-fs/pkg/global"
)

var handler *webdav.Handler

// webdav 路由
func Route(dav *echo.Group) {
    handler = &webdav.Handler{
        Prefix:     "/dav",
        FileSystem: webdav.Dir(global.Conf.Webdav.Path),
        LockSystem: webdav.NewMemLS(),
    }

    dav.Use(WebDAVAuth())
    dav.Any("/*", ServeWebDAV)
    dav.Any("", ServeWebDAV)
    dav.Add("PROPFIND", "/*", ServeWebDAV)
    dav.Add("PROPFIND", "", ServeWebDAV)
    dav.Add("MKCOL", "/*", ServeWebDAV)
    dav.Add("LOCK", "/*", ServeWebDAV)
    dav.Add("UNLOCK", "/*", ServeWebDAV)
    dav.Add("PROPPATCH", "/*", ServeWebDAV)
    dav.Add("COPY", "/*", ServeWebDAV)
    dav.Add("MOVE", "/*", ServeWebDAV)
}

// ServeWebDAV
func ServeWebDAV(ctx echo.Context) error {
    req := ctx.Request()
    w := ctx.Response()

    handler.ServeHTTP(w, req)

    return nil
}

func WebDAVAuth() echo.MiddlewareFunc {
    return func(next echo.HandlerFunc) echo.HandlerFunc {
        return func(ctx echo.Context) error {
            req := ctx.Request()
            w := ctx.Response()

            // 获取用户名/密码
            username, password, ok := req.BasicAuth()

            if !ok {
                w.Header().Set("WWW-Authenticate", `Basic realm="Restricted"`)
                w.WriteHeader(http.StatusUnauthorized)
                return nil
            }

            userPassword := global.Conf.Webdav.GetUserPassword(username)

            // 验证用户名/密码
            if userPassword == "" || !utils.PasswordCheck(password, userPassword) {
                http.Error(w, "WebDAV: need authorized!", http.StatusUnauthorized)
                return nil
            }

            return next(ctx)
        }
    }
}
