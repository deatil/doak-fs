package boot

import (
    "fmt"
    "time"
    "flag"
    "strings"
    "net/http"

    "github.com/labstack/echo/v4"
    "github.com/labstack/echo/v4/middleware"
    "github.com/labstack/gommon/log"
    "github.com/labstack/gommon/color"

    "github.com/deatil/doak-fs/pkg/utils"
    "github.com/deatil/doak-fs/pkg/global"
    "github.com/deatil/doak-fs/pkg/logger"
    "github.com/deatil/doak-fs/pkg/session"
    "github.com/deatil/doak-fs/pkg/response"
    "github.com/deatil/doak-fs/pkg/template"

    "github.com/deatil/doak-fs/app/view"
    "github.com/deatil/doak-fs/app/route"
    "github.com/deatil/doak-fs/app/webdav"
    "github.com/deatil/doak-fs/app/resources"
)

const (
    website = "https://github.com/deatil/doak-fs"
    banner = `
    .___            __               _____
  __| _/_________  |  | __         _/ ____\______
 / __ |/  _ \__  \ |  |/ /  ______ \   __\/  ___/
/ /_/ (  <_> ) __ \|    <  /_____/  |  |  \___ \
\____ |\____(____  /__|_ \          |__| /____  >
     \/          \/     \/                    \/  %s
Doak filesystem base on echo %s
%s
_____________________________________________________

`
)

var makePass string

// 初始化
func initServer() {
    // 系统启动参数
    config := flag.String("config", "", "config file")
    view   := flag.String("view", "", "view path")
    pass   := flag.String("pass", "", "make Pass")
    flag.Parse()

    global.ConfigFile = *config
    global.ViewPath = *view

    makePass = *pass

    // 只使用打包文件
    global.IsOnlyEmbed = true

    initConfig()

    initTime()

    initFs()
}

// 运行
func Start() {
    initServer()

    if makePass != "" {
        newPass := utils.PasswordHash(makePass)
        fmt.Println("new password: " + newPass)
        return
    }

    runServer()
}

// 运行
func runServer() {
    // 初始化 echo
    e := echo.New()

    // 设置日志信息
    logger.SetLoggerFile(global.Conf.App.LogFile)
    logger.SetLoggerLevel(global.Conf.App.LogLevel)

    // 隐藏默认
    e.HideBanner = true
    e.HidePort = true

    // 自定义错误处理
    e.HTTPErrorHandler = httpErrorHandler

    // 调试状态
    debug := global.Conf.App.Debug

    // 移除 url 结尾 /
    e.Pre(middleware.RemoveTrailingSlash())

    // 超时处理
    e.Use(middleware.TimeoutWithConfig(middleware.TimeoutConfig{
        Skipper: middleware.DefaultSkipper,
        OnTimeoutRouteErrorHandler: func(err error, ctx echo.Context) {
            ctx.Logger().Error(err)
        },
        Timeout: 50 * time.Second,
    }))

    // 拦截报错
    recoverLength := 8 << 10 // 8 KB
    e.Use(middleware.RecoverWithConfig(middleware.RecoverConfig{
        Skipper:           middleware.DefaultSkipper,
        StackSize:         recoverLength,
        DisableStackAll:   false,
        DisablePrintStack: false,
        LogLevel:          log.ERROR,
        LogErrorFunc:      func(ctx echo.Context, err error, stack []byte) error {
            // 记录日志
            logger.Logger().WithFields(logger.Fields{
                "Error": err.Error(),
                "Stack": string(stack),
            }).Error("PANIC RECOVER")

            // 打印日志
            if debug {
                msg := fmt.Sprintf("[PANIC RECOVER] %v %s\n", err, stack[:recoverLength])

                ctx.Logger().Error(msg)
            }

            return err
        },
    }))

    // 跳转到 www 开头
    // e.Use(middleware.WWWRedirect())

    // 加密
    e.Use(middleware.SecureWithConfig(middleware.SecureConfig{
        Skipper:            middleware.DefaultSkipper,
        XSSProtection:      "1; mode=block",
        ContentTypeNosniff: "nosniff",
        XFrameOptions:      "SAMEORIGIN",
        HSTSPreloadEnabled: false,
    }))

    // CSRF
    e.Use(middleware.CSRFWithConfig(middleware.CSRFConfig{
        Skipper:        func(ctx echo.Context) bool {
            path := ctx.Request().URL.String()
            if strings.HasPrefix(path, "/dav") {
                return true
            }

            return false
        },
        TokenLength:    global.Conf.Server.CSRFTokenLength,
        TokenLookup:    "cookie:" + global.Conf.Server.CSRFCookieName,
        ContextKey:     global.Conf.Server.CSRFContextKey,
        CookieName:     global.Conf.Server.CSRFCookieName,
        CookiePath:     global.Conf.Server.CSRFCookiePath,
        CookieMaxAge:   global.Conf.Server.CSRFCookieMaxAge,
        // SameSiteDefaultMode | SameSiteLaxMode | SameSiteStrictMode | SameSiteNoneMode
        CookieSameSite: http.SameSiteDefaultMode,
        ErrorHandler:   func(err error, ctx echo.Context) error {
            logger.Logger().WithFields(logger.Fields{
                "Error": err.Error(),
            }).Error("CSRF")

            return err
        },
    }))

    // Gzip
    e.Use(middleware.GzipWithConfig(middleware.GzipConfig{
        Skipper: middleware.DefaultSkipper,
        Level:   5,
    }))

    // Decompress
    e.Use(middleware.Decompress())

    // 设置日志
    // e.Use(middleware.Logger())
    e.Use(middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
        LogURI:    true,
        LogStatus: true,
        LogValuesFunc: func(ctx echo.Context, values middleware.RequestLoggerValues) error {
            req := ctx.Request()

            if debug {
                if values.Latency > time.Minute {
                    values.Latency = values.Latency - values.Latency%time.Second
                }

                // 显示请求信息
                requestData := fmt.Sprintf("[doak-fs] %s | %3d | %15s | %-15s | %-7s | %s | %#v",
                    values.StartTime.Format("2006-01-02 15:04:05"),
                    values.Status,
                    values.Latency,
                    ctx.RealIP(),
                    req.Method,
                    values.URI,
                    values.Error,
                )

                fmt.Println(requestData)
            }

            // 记录报错信息
            if values.Error != nil {
                logger.Logger().WithFields(logger.Fields{
                    "Time":      values.StartTime.Format("2006-01-02 15:04:05"),
                    "Method":    req.Method,
                    "URI":       values.URI,
                    "Status":    values.Status,
                    "RemoteIP":  ctx.RealIP(),
                    "Latency":   fmt.Sprintf("%v", values.Latency),
                    "UserAgent": values.UserAgent,
                    "Error":     values.Error.Error(),
                }).Error("request")
            }

            return nil
        },
    }))

    // 设置 seesion
    session.SessionMiddleware(e)

    // 设置模板
    renderer := template.NewTemplate()
    renderer.SetDebug(debug)
    renderer.AddFuncs(view.ViewFuncs())

    if global.ViewPath != "" && !global.IsOnlyEmbed {
        renderer.SetUseEmbed(false)
        renderer.AddDirectory(global.ViewPath)
    } else {
        renderer.SetUseEmbed(true)
        renderer.SetEmbedReadFileFunc(resources.ReadViewFile)
    }

    e.Renderer = renderer

    // 静态文件
    assetHandler := http.FileServer(resources.StaticAssets())
    e.GET("/static/*", echo.WrapHandler(http.StripPrefix("/static/", assetHandler)))

    // webdav
    webdav.Route(e.Group("/dav"))

    // 路由
    route.Route(e)

    // 未发现路由
    e.RouteNotFound("/*", func(ctx echo.Context) error {
        if ctx.Request().Method == "POST" {
            return response.ReturnErrorJson(ctx, "not found")
        }

        return ctx.String(http.StatusNotFound, "not found")
    })

    // 显示信息
    showBanner()

    // 设置端口
    e.Logger.Fatal(e.Start(global.Conf.Server.Address))
}

// 自定义错误
func httpErrorHandler(err error, ctx echo.Context) {
    code := http.StatusInternalServerError
    if he, ok := err.(*echo.HTTPError); ok {
        code = he.Code
    }

    ctx.Logger().Error(err)

    errMsg := "Server Error!"
    if global.Conf.App.Debug {
        errMsg = err.Error()
    }

    var repErr error

    if ctx.Request().Method == "POST" {
        repErr = response.ReturnErrorJson(ctx, errMsg)
    } else {
        repErr = ctx.String(code, errMsg)
    }

    // 输出字符
    if repErr != nil {
        ctx.Logger().Error(repErr)
    }
}

// 控制台显示信息
func showBanner() {
    colorer := color.New()

    colorer.Printf(
        banner,
        colorer.Red("v" + global.Conf.App.Version),
        colorer.Green("v" + echo.Version),
        colorer.Blue(website),
    )

    if global.Conf.App.Debug {
        colorer.Printf("https server started on %s\n", colorer.Green(global.Conf.Server.Address))
    }
}
