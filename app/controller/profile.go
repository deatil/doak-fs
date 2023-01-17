package controller

import (
    "github.com/labstack/echo/v4"

    "github.com/deatil/doak-fs/pkg/utils"
    "github.com/deatil/doak-fs/pkg/config"
    "github.com/deatil/doak-fs/pkg/global"
    "github.com/deatil/doak-fs/pkg/session"
    "github.com/deatil/doak-fs/pkg/response"

    "github.com/deatil/doak-fs/app/model"
)

/**
 * 我的信息
 *
 * @create 2023-1-17
 * @author deatil
 */
type Profile struct{
    Base
}

// 更改密码页面
func (this *Profile) Password(ctx echo.Context) error {
    username := ctx.Get("username")

    return response.Render(ctx, "profile_password.html", map[string]any{
        "username": username,
    })
}

// 更改密码保存
func (this *Profile) PasswordSave(ctx echo.Context) error {
    oldPass := ctx.FormValue("old_pass")
    newPass := ctx.FormValue("new_pass")
    newPassCheck := ctx.FormValue("new_pass_check")

    if oldPass == "" {
        return response.ReturnErrorJson(ctx, "旧密码不能为空")
    }
    if newPass == "" {
        return response.ReturnErrorJson(ctx, "新密码不能为空")
    }
    if newPassCheck == "" {
        return response.ReturnErrorJson(ctx, "确认密码不能为空")
    }

    if newPass != newPassCheck {
        return response.ReturnErrorJson(ctx, "确认密码不一致")
    }

    username := session.Get(ctx, "userid").(string)

    userPassword := model.GetUserPassword(username)
    if !utils.PasswordCheck(oldPass, userPassword) {
        return response.ReturnErrorJson(ctx, "旧密码错误")
    }

    // 新密码
    newMakePass := utils.PasswordHash(newPass)

    // 更改密码
    global.Conf.User = global.Conf.User.UpdatePassword(username, newMakePass)

    // 更改配置信息
    if global.ConfigFile != "" && !global.IsOnlyEmbed {
        // 读取配置
        err := config.WriteConfig(global.ConfigFile, global.Conf)
        if err != nil {
            return response.ReturnErrorJson(ctx, "更改密码失败")
        }
    }

    return response.ReturnSuccessJson(ctx, "更改密码成功", "")
}
