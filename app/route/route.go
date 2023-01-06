package route

import (
    "github.com/labstack/echo/v4"

    "github.com/deatil/doak-fs/app/controller"
    "github.com/deatil/doak-fs/app/middleware"
)

// 路由
func Route(e *echo.Echo) {
    indexController := new(controller.Index)
    e.GET("/", indexController.Index, middleware.Auth())

    // 登录相关
    authController := new(controller.Auth)
    e.GET("/auth/captcha", authController.Captcha)
    e.GET("/auth/login", authController.Login)
    e.POST("/auth/login", authController.LoginCheck)
    e.GET("/auth/logout", authController.Logout)

    // 文件管理
    fileGroup := e.Group("/file", middleware.Auth())
    {
        fileController := new(controller.File)
        fileGroup.GET("/index", fileController.Index)
        fileGroup.GET("/info", fileController.Info)
        fileGroup.POST("/delete", fileController.Delete)

        fileGroup.POST("/create-file", fileController.CreateFile)
        fileGroup.POST("/rename-file", fileController.RenameFile)
        fileGroup.GET("/update-file", fileController.UpdateFile)
        fileGroup.POST("/update-file", fileController.UpdateFileSave)
        fileGroup.POST("/upload-file", fileController.UploadFile)
        fileGroup.GET("/download-file", fileController.DownloadFile)
        fileGroup.POST("/move-file", fileController.MoveFile)

        fileGroup.POST("/make-dir", fileController.MakeDir)
        fileGroup.POST("/move-dir", fileController.MoveDir)
        fileGroup.POST("/rename-dir", fileController.RenameDir)
    }
}

