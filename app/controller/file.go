package controller

import (
    "io"
    "os"

    "github.com/labstack/echo/v4"

    "github.com/deatil/doak-fs/pkg/fs"
    "github.com/deatil/doak-fs/pkg/global"
    "github.com/deatil/doak-fs/pkg/response"
)

/**
 * 文件管理
 *
 * @create 2023-1-2
 * @author deatil
 */
type File struct{
    Base
}

// 首页
func (this *File) Index(ctx echo.Context) error {
    path := ctx.QueryParam("path")

    // 根目录
    rootPath := global.Conf.File.Path
    filePath, _ := fs.Filesystem.Realpath(rootPath)

    if path != "" {
        filePath = fs.JoinPath(filePath, path)
    }

    list := fs.Ls(filePath)
    name := fs.Basename(filePath)

    parentPath := ""

    if path == "" || path == "/" {
        path = "/"
        name = "/"
    }

    if path != "/" {
        parentPath = fs.Filesystem.Dirname(path)
        parentPath = fs.Filesystem.ToSlash(parentPath)
    }

    username := ctx.Get("username")

    return response.Render(ctx, "file_index.html", map[string]any{
        "path": path,
        "name": name,
        "filePath": filePath,
        "parentPath": parentPath,
        "list": list,

        "username": username,
    })
}

// 详情
func (this *File) Info(ctx echo.Context) error {
    file := ctx.FormValue("file")
    if file == "" {
        return response.ReturnErrorJson(ctx, "文件不能为空")
    }

    rootPath := global.Conf.File.Path
    filePath, _ := fs.Filesystem.Realpath(rootPath)

    filePath = fs.JoinPath(filePath, file)

    data := fs.Read(filePath)

    return response.Render(ctx, "file_info.html", map[string]any{
        "data": data,
        "file": file,
    })
}

// 删除文件
func (this *File) Delete(ctx echo.Context) error {
    file := ctx.FormValue("file")
    if file == "" {
        return response.ReturnErrorJson(ctx, "文件不能为空")
    }

    rootPath := global.Conf.File.Path
    filePath, _ := fs.Filesystem.Realpath(rootPath)

    filePath = fs.JoinPath(filePath, file)

    if err := fs.Delete(filePath); err != nil {
        return response.ReturnErrorJson(ctx, "删除文件失败")
    }

    return response.ReturnSuccessJson(ctx, "删除文件成功", "")
}

// 创建文件
func (this *File) CreateFile(ctx echo.Context) error {
    return response.ReturnSuccessJson(ctx, "创建文件成功", "")
}

// 重命名文件
func (this *File) RenameFile(ctx echo.Context) error {
    return response.ReturnSuccessJson(ctx, "重命名文件成功", "")
}

// 更新文件
func (this *File) UpdateFile(ctx echo.Context) error {
    file := ctx.FormValue("file")
    if file == "" {
        return response.String(ctx, "文件不能为空")
    }

    rootPath := global.Conf.File.Path
    filePath, _ := fs.Filesystem.Realpath(rootPath)

    filePath = fs.JoinPath(filePath, file)

    if !fs.IsFile(filePath) {
        return response.String(ctx, "打开的不是文件")
    }

    data, err := fs.Get(filePath)
    if err != nil {
        return response.String(ctx, "打开文件失败")
    }

    ext := fs.Filesystem.Extension(filePath)

    return response.Render(ctx, "file_update_file.html", map[string]any{
        "data": data,
        "file": file,
        "ext": ext,
    })
}

// 更新文件保存
func (this *File) UpdateFileSave(ctx echo.Context) error {
    file := ctx.FormValue("file")
    if file == "" {
        return response.ReturnErrorJson(ctx, "文件不能为空")
    }

    data := ctx.FormValue("data")
    if data == "" {
        return response.ReturnErrorJson(ctx, "文件内容不能为空")
    }

    rootPath := global.Conf.File.Path
    filePath, _ := fs.Filesystem.Realpath(rootPath)

    filePath = fs.JoinPath(filePath, file)

    if !fs.IsFile(filePath) {
        return response.ReturnErrorJson(ctx, "要更新的不是文件")
    }

    if err := fs.Put(filePath, data); err != nil {
        return response.ReturnErrorJson(ctx, "更新文件失败")
    }

    return response.ReturnSuccessJson(ctx, "更新文件成功", "")
}

// 上传文件
func (this *File) UploadFile(ctx echo.Context) error {
    file, err := ctx.FormFile("file")
    if err != nil {
        return err
    }

    src, err := file.Open()
    if err != nil {
        return err
    }
    defer src.Close()

    // 路径
    path := ctx.QueryParam("path")
    if path == "" {
        return response.ReturnErrorJson(ctx, "保存路径不能为空")
    }

    // 根目录
    rootPath := global.Conf.File.Path
    filePath, _ := fs.Filesystem.Realpath(rootPath)

    filename := fs.JoinPath(filePath, path, file.Filename)

    // 创建文件
    dst, err := os.Create(filename)
    if err != nil {
        return err
    }
    defer dst.Close()

    // 保存
    if _, err = io.Copy(dst, src); err != nil {
        return err
    }

    return response.ReturnSuccessJson(ctx, "上传文件成功", "")
}

// 下载文件
func (this *File) DownloadFile(ctx echo.Context) error {
    file := ctx.QueryParam("file")
    if file == "" {
        return response.String(ctx, "文件不能为空")
    }

    rootPath := global.Conf.File.Path
    filePath, _ := fs.Filesystem.Realpath(rootPath)

    filePath = fs.JoinPath(filePath, file)

    if !fs.IsFile(filePath) {
        return response.String(ctx, "打开的不是文件")
    }

    basename := fs.Filesystem.Basename(filePath)

    return response.Attachment(ctx, filePath, basename)
}

// 移动文件
func (this *File) MoveFile(ctx echo.Context) error {
    return response.ReturnSuccessJson(ctx, "移动文件成功", "")
}

// 创建文件夹
func (this *File) MakeDir(ctx echo.Context) error {
    return response.ReturnSuccessJson(ctx, "创建文件夹成功", "")
}

// 移动文件夹
func (this *File) MoveDir(ctx echo.Context) error {
    return response.ReturnSuccessJson(ctx, "移动文件夹成功", "")
}

// 重命名文件夹
func (this *File) RenameDir(ctx echo.Context) error {
    return response.ReturnSuccessJson(ctx, "重命名文件夹成功", "")
}

