## doak-fs 文件管理

文件管理工具，web界面，使用 web 框架 `echo`，打包配置文件、静态文件和模板文件，生成一个执行文件部署方便，也可以设置参数自定义配置文件和模板目录


### 项目介绍

*  使用 `go` 开发的文件管理工具
*  web 框架使用 `go` 框架 `echo`
*  模板库使用 `pongo2` 库，语法接近 `python` 的 `django` 框架
*  默认打包配置文件、静态文件和模板文件。可更改 `app/boot/boot.go` 文件内 `global.IsOnlyEmbed` 为 `false` 自定义配置文件和模板文件
*  生成一个文件即可部署
*  添加 WebDAV 支持


### 截图预览

![login](https://user-images.githubusercontent.com/24578855/219612410-b6994879-33d1-49d4-818e-d6d3be9fe50d.png)
![index](https://user-images.githubusercontent.com/24578855/219612392-f1555a54-ea09-441a-a1c2-eaf17b2b37d6.png)
![file](https://user-images.githubusercontent.com/24578855/219553564-dbd3dea2-df6e-4d0d-b6ba-ac0a94ec12c9.png)

### 使用方法

1. 下载

```cmd
git clone github.com/deatil/doak-fs
```

2. 编译运行

```cmd
go run main.go
```

自定义配置文件
```cmd
go run main.go --config=config.toml
```

使用模板位置
```cmd
go run main.go --view=template
```

3. 登录账号: `admin` / `123456`, WebDAV 账号: `webnav` / `123456`


### 特别鸣谢

感谢以下的项目,排名不分先后

 - github.com/labstack/echo

 - github.com/flosch/pongo2

 - github.com/jinzhu/now

 - github.com/deatil/lakego-filesystem

 - github.com/steambap/captcha


### 开源协议

*  `doak-fs` 遵循 `Apache2` 开源协议发布，在保留本软件版权的情况下提供个人及商业免费使用。


### 版权

*  该系统所属版权归 deatil(https://github.com/deatil) 所有。
