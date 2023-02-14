package config

import (
    "os"
    "fmt"
    "strings"

    "github.com/pelletier/go-toml/v2"
)

type Conf struct {
    App     `toml:"app"`
    Server  `toml:"server"`
    File    `toml:"file"`
    User    `toml:"user"`
    Session `toml:"session"`
}

type App struct {
    Appname  string `toml:"app_name"`
    Version  string `toml:"version"`
    Debug    bool   `toml:"debug"`
    TimeZone string `toml:"time_zone"`
    LogFile  string `toml:"log_file"`
    LogLevel string `toml:"log_level"`
    Assets   string `toml:"assets"`
}

type Server struct {
    Address          string `toml:"address"`
    CSRFTokenLength  uint8  `toml:"csrf_token_length"`
    CSRFContextKey   string `toml:"csrf_context_key"`
    CSRFCookieName   string `toml:"csrf_cookie_name"`
    CSRFCookiePath   string `toml:"csrf_cookie_path"`
    CSRFCookieMaxAge int    `toml:"csrf_cookie_maxage"`
}

type File struct {
    Path   string `toml:"path"`
    Driver string `toml:"driver"`
}

type User struct {
    Names []string `toml:"names"`
}

// 更改密码
func (this User) GetUsers() map[string]string {
    users := make(map[string]string)

    for _, name := range this.Names {
        newName := strings.SplitN(name, ":", 2)
        users[newName[0]] = newName[1]
    }

    return users
}

// 账号密码
func (this User) GetUserPassword(name string) string {
    users := this.GetUsers()

    if password, ok := users[name]; ok {
        return password
    }

    return ""
}

// 更改密码
func (this User) UpdatePassword(name string, pass string) User {
    users := this.GetUsers()

    _, ok := users[name]
    if !ok {
        return this
    }

    users[name] = pass

    this.Names = make([]string, 0)
    for name, pass := range users {
        this.Names = append(this.Names, fmt.Sprintf("%s:%s", name, pass))
    }

    return this
}

type Session struct {
    Secret   string `toml:"secret"`
    Key      string `toml:"key"`
    Path     string `toml:"path"`
    MaxAge   int    `toml:"max_age"`
    HttpOnly bool   `toml:"http_only"`
}

// 读取配置文件
func ReadConfig(file string) (Conf, error) {
    var conf Conf

    fs, err := os.Open(file)
    if err != nil {
        return conf, err
    }
    defer fs.Close()

    err = toml.NewDecoder(fs).Decode(&conf)
    if err != nil {
        return conf, err
    }

    return conf, nil
}

// 读取配置文件
func ReadConfigByte(data []byte) (Conf, error) {
    var conf Conf

    err := toml.Unmarshal([]byte(data), &conf)
    if err != nil {
        return conf, err
    }

    return conf, nil
}

// 写入配置文件
func WriteConfig(file string, cfg Conf) error {
    b, err := toml.Marshal(cfg)
    if err != nil {
        return err
    }

    err = os.WriteFile(file, b, 0666)
    if err != nil {
        return err
    }

    return nil
}

// 写入配置文件
func WriteConfigToByte(cfg Conf) ([]byte, error) {
    return toml.Marshal(cfg)
}
