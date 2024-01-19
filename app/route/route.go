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

    // 我的信息
    profileGroup := e.Group("/profile", middleware.Auth())
    {
        profileController := new(controller.Profile)
        profileGroup.GET("/password", profileController.Password)
        profileGroup.POST("/password", profileController.PasswordSave)

        profileGroup.GET("/webdav", profileController.Webdav)
        profileGroup.POST("/webdav", profileController.WebdavSave)
    }

    // 文件管理
    fileGroup := e.Group("/file", middleware.Auth())
    {
        fileController := new(controller.File)
        fileGroup.GET("/index", fileController.Index)
        fileGroup.GET("/info", fileController.Info)
        fileGroup.POST("/delete", fileController.Delete)
        fileGroup.POST("/rename", fileController.Rename)
        fileGroup.GET("/move", fileController.Move)
        fileGroup.POST("/move", fileController.MoveSave)
        fileGroup.GET("/copy", fileController.Copy)
        fileGroup.POST("/copy", fileController.CopySave)
        fileGroup.GET("/upload", fileController.Upload)
        fileGroup.POST("/upload", fileController.UploadSave)

        fileGroup.POST("/create-file", fileController.CreateFile)
        fileGroup.GET("/update-file", fileController.UpdateFile)
        fileGroup.POST("/update-file", fileController.UpdateFileSave)

        fileGroup.POST("/create-dir", fileController.CreateDir)

        // 本地相关
        fileGroup.GET("/download-file", fileController.DownloadFile)
        fileGroup.GET("/preview-file", fileController.PreviewFile)
    }
}

