package controller

import (
    "github.com/labstack/echo/v4"

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

    dirList := global.Fs.LsDir(path)
    fileList := global.Fs.LsFile(path)

    list := append(dirList, fileList...)

    name := global.Fs.Basename(path)

    parentPath := global.Fs.ParentPath(path)

    if path == "" || path == "/" {
        path = ""
        name = "/"
    }

    username := ctx.Get("username")

    return response.Render(ctx, "file_index.html", map[string]any{
        "path": path,
        "name": name,
        "parentPath": parentPath,
        "list": list,

        "username": username,
    })
}

// 详情
func (this *File) Info(ctx echo.Context) error {
    file := ctx.FormValue("file")
    if file == "" {
        return response.String(ctx, "文件不能为空")
    }

    data := global.Fs.Read(file)

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

    if err := global.Fs.Delete(file); err != nil {
        return response.ReturnErrorJson(ctx, "删除失败")
    }

    return response.ReturnSuccessJson(ctx, "删除成功", "")
}

// 重命名
func (this *File) Rename(ctx echo.Context) error {
    oldName := ctx.FormValue("old_name")
    if oldName == "" {
        return response.ReturnErrorJson(ctx, "旧名称不能为空")
    }

    newName := ctx.FormValue("new_name")
    if newName == "" {
        return response.ReturnErrorJson(ctx, "新名称不能为空")
    }

    if err := global.Fs.Rename(oldName, newName); err != nil {
        return response.ReturnErrorJson(ctx, "重命名失败")
    }

    return response.ReturnSuccessJson(ctx, "重命名成功", "")
}

// 移动
func (this *File) Move(ctx echo.Context) error {
    old := ctx.FormValue("old")
    if old == "" {
        return response.String(ctx, "访问错误")
    }

    path := ctx.QueryParam("path")

    list := global.Fs.LsDir(path)
    parentPath := global.Fs.ParentPath(path)

    if path == "" || path == "/" {
        path = "/"
    }

    return response.Render(ctx, "file_move.html", map[string]any{
        "old": old,
        "path": path,
        "parentPath": parentPath,
        "list": list,
    })
}

// 移动保存
func (this *File) MoveSave(ctx echo.Context) error {
    oldName := ctx.FormValue("old_name")
    if oldName == "" {
        return response.ReturnErrorJson(ctx, "旧名称不能为空")
    }

    newName := ctx.FormValue("new_name")
    if newName == "" {
        return response.ReturnErrorJson(ctx, "新名称不能为空")
    }

    if err := global.Fs.Move(oldName, newName); err != nil {
        return response.ReturnErrorJson(ctx, "移动失败。原因: " + err.Error())
    }

    return response.ReturnSuccessJson(ctx, "移动成功", "")
}

// 复制
func (this *File) Copy(ctx echo.Context) error {
    old := ctx.FormValue("old")
    if old == "" {
        return response.String(ctx, "访问错误")
    }

    path := ctx.QueryParam("path")

    list := global.Fs.LsDir(path)
    parentPath := global.Fs.ParentPath(path)

    if path == "" || path == "/" {
        path = "/"
    }

    return response.Render(ctx, "file_copy.html", map[string]any{
        "old": old,
        "path": path,
        "parentPath": parentPath,
        "list": list,
    })
}

// 复制保存
func (this *File) CopySave(ctx echo.Context) error {
    oldName := ctx.FormValue("old_name")
    if oldName == "" {
        return response.ReturnErrorJson(ctx, "旧名称不能为空")
    }

    newName := ctx.FormValue("new_name")
    if newName == "" {
        return response.ReturnErrorJson(ctx, "新名称不能为空")
    }

    if err := global.Fs.Copy(oldName, newName); err != nil {
        return response.ReturnErrorJson(ctx, "复制失败。原因: " + err.Error())
    }

    return response.ReturnSuccessJson(ctx, "复制成功", "")
}

// 创建文件
func (this *File) CreateFile(ctx echo.Context) error {
    file := ctx.FormValue("file")
    if file == "" {
        return response.ReturnErrorJson(ctx, "文件不能为空")
    }

    if err := global.Fs.CreateFile(file); err != nil {
        return response.ReturnErrorJson(ctx, "创建文件失败")
    }

    return response.ReturnSuccessJson(ctx, "创建文件成功", "")
}

// 更新文件
func (this *File) UpdateFile(ctx echo.Context) error {
    file := ctx.FormValue("file")
    if file == "" {
        return response.String(ctx, "文件不能为空")
    }

    info := global.Fs.Read(file)
    if info["size"].(int64) > 15728640 {
        return response.String(ctx, "文件太大不支持打开")
    }

    data, err := global.Fs.Get(file)
    if err != nil {
        return response.String(ctx, err.Error())
    }

    ext := global.Fs.Extension(file)

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

    if err := global.Fs.Put(file, data); err != nil {
        return response.ReturnErrorJson(ctx, err.Error())
    }

    return response.ReturnSuccessJson(ctx, "更新文件成功", "")
}

// 上传文件
func (this *File) Upload(ctx echo.Context) error {
    path := ctx.FormValue("path")
    if path == "" {
        path = "/"
    }

    return response.Render(ctx, "file_upload.html", map[string]any{
        "path": path,
    })
}

// 上传文件保存
func (this *File) UploadSave(ctx echo.Context) error {
    file, err := ctx.FormFile("file")
    if err != nil {
        return err
    }

    src, err := file.Open()
    if err != nil {
        return response.ReturnErrorJson(ctx, "上传文件错误")
    }
    defer src.Close()

    // 路径
    path := ctx.FormValue("path")
    if path == "" {
        return response.ReturnErrorJson(ctx, "保存路径不能为空")
    }

    // 保存
    if err = global.Fs.Upload(src, path, file.Filename); err != nil {
        return response.ReturnErrorJson(ctx, err.Error())
    }

    return response.ReturnSuccessJson(ctx, "上传文件成功", "")
}

// 创建文件夹
func (this *File) CreateDir(ctx echo.Context) error {
    dir := ctx.FormValue("dir")
    if dir == "" {
        return response.ReturnErrorJson(ctx, "文件夹不能为空")
    }

    if err := global.Fs.CreateDir(dir); err != nil {
        return response.ReturnErrorJson(ctx, "创建文件夹失败")
    }

    return response.ReturnSuccessJson(ctx, "创建文件夹成功", "")
}

// ============== 本地相关 ==============

// 下载文件
func (this *File) DownloadFile(ctx echo.Context) error {
    file := ctx.QueryParam("file")
    if file == "" {
        return response.String(ctx, "文件不能为空")
    }

    filePath, err := global.Fs.FormatFile(file)
    if err != nil {
        return response.String(ctx, "文件错误")
    }

    basename := global.Fs.Basename(filePath)

    return response.Attachment(ctx, filePath, basename)
}

// 预览文件
func (this *File) PreviewFile(ctx echo.Context) error {
    file := ctx.QueryParam("file")
    if file == "" {
        return response.String(ctx, "文件不能为空")
    }

    data := global.Fs.Read(file)
    dataType := data["type"].(string)

    if dataType != "image" &&
        dataType != "audio" &&
        dataType != "video" {
        return response.String(ctx, "文件不存在")
    }

    filePath, err := global.Fs.FormatFile(file)
    if err != nil {
        return response.String(ctx, "文件错误")
    }

    return response.File(ctx, filePath)
}

