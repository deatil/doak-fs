package controller

import (
    "strings"
    "github.com/labstack/echo/v4"
    "github.com/steambap/captcha"
    "golang.org/x/image/font/gofont/goregular"

    "github.com/deatil/doak-fs/pkg/utils"
    "github.com/deatil/doak-fs/pkg/session"
    "github.com/deatil/doak-fs/pkg/response"

    "github.com/deatil/doak-fs/app/model"
)

/**
 * 登录
 *
 * @create 2022-12-30
 * @author deatil
 */
type Auth struct{
    Base
}

// 验证码
func (this *Auth) Captcha(ctx echo.Context) error {
    captcha.LoadFont(goregular.TTF)

    img, err := captcha.New(125, 46, func(options *captcha.Options) {
        options.FontScale = 0.8
    })

    if err != nil {
        return response.String(ctx, "error")
    }

    img.WriteImage(ctx.Response())

    // 存储验证码
    session.Set(ctx, "captchaid", img.Text)

    return nil
}

// 登录页面
func (this *Auth) Login(ctx echo.Context) error {
    return response.Render(ctx, "auth_login.html", map[string]any{})
}

// 检测登录
func (this *Auth) LoginCheck(ctx echo.Context) error {
    username := ctx.FormValue("username")
    password := ctx.FormValue("password")
    captchaData := ctx.FormValue("captcha")

    if username == "" {
        return response.ReturnErrorJson(ctx, "账号不能为空")
    }
    if password == "" {
        return response.ReturnErrorJson(ctx, "密码不能为空")
    }
    if captchaData == "" {
        return response.ReturnErrorJson(ctx, "验证码不能为空")
    }

    // 验证码ID
    captchaid := session.Get(ctx, "captchaid")

    // 清除数据
    session.Delete(ctx, "captchaid")

    if strings.ToLower(captchaid.(string)) != strings.ToLower(captchaData) {
        return response.ReturnErrorJson(ctx, "验证码错误")
    }

    userPassword := model.GetUserPassword(username)
    if userPassword == "" || !utils.PasswordCheck(password, userPassword) {
        return response.ReturnErrorJson(ctx, "账号不存在或者密码错误")
    }

    // 保存状态
    session.Set(ctx, "userid", username)

    return response.ReturnSuccessJson(ctx, "登录成功", "")
}

// 退出
func (this *Auth) Logout(ctx echo.Context) error {
    userid := session.Get(ctx, "userid")
    if userid == nil {
        return response.Redirect(ctx, "/auth/login")
    }

    // 删除登录数据
    session.Delete(ctx, "userid")

    return response.Redirect(ctx, "/auth/login")
}

