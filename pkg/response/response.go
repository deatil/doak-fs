package response

import (
    "io"
    "strings"
    "net/http"
    "encoding/json"

    "github.com/labstack/echo/v4"
)

/**
 * json 返回数据
 *
 * @create 2022-12-30
 * @author deatil
 */
type JSONResult struct {
    Code    int    `json:"code"`
    Message string `json:"message"`
    Data    any    `json:"data"`
}

// 返回 json
func ReturnJson(
    ctx  echo.Context,
    code int,
    msg  string,
    data any,
) error {
    return JSON(ctx, JSONResult{
        Code:    code,
        Message: msg,
        Data:    data,
    })
}

// 返回成功 json
func ReturnSuccessJson(ctx echo.Context, msg string, data any) error {
    return ReturnJson(ctx, 0, msg, data)
}

// 返回错误 json
func ReturnErrorJson(ctx echo.Context, msg string) error {
    return ReturnJson(ctx, 1, msg, "")
}

// 响应模板
func Render(ctx echo.Context, name string, data any) error {
    return ctx.Render(http.StatusOK, name, data)
}

// 响应字符输出
func HTML(ctx echo.Context, html string) error {
    return ctx.HTML(http.StatusOK, html)
}

// 响应字符
func String(ctx echo.Context, str string) error {
    return ctx.String(http.StatusOK, str)
}

// 响应 JSON
func JSON(ctx echo.Context, i any) error {
    return ctx.JSON(http.StatusOK, i)
}

// json
func JSONString(ctx echo.Context, i any) error {
    ctx.Response().Header().Set(echo.HeaderContentType, echo.MIMEApplicationJSONCharsetUTF8)
    ctx.Response().WriteHeader(http.StatusOK)

    return json.NewEncoder(ctx.Response()).Encode(i)
}

// 响应 JSONP
func JSONP(ctx echo.Context, callback string, i any) error {
    return ctx.JSONP(http.StatusOK, callback, i)
}

// 响应 XML
func XML(ctx echo.Context, i any) error {
    return ctx.XML(http.StatusOK, i)
}

// 响应数据流
func Stream(ctx echo.Context, contentType string, r io.Reader) error {
    return ctx.Stream(http.StatusOK, contentType, r)
}

// 下载文件
func File(ctx echo.Context, file string) error {
    return ctx.File(file)
}

// 下载附件
func Attachment(ctx echo.Context, file string, name string) error {
    return ctx.Attachment(file, name)
}

// 跳转
func Redirect(ctx echo.Context, url string) error {
    return ctx.Redirect(http.StatusSeeOther, url)
}

// IsAjax
func IsAjax(ctx echo.Context) bool {
    header := ctx.Request().Header

    return strings.EqualFold(header.Get("X-Requested-With"), "XMLHttpRequest")
}
