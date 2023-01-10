package logger

import (
    "os"
    "sync"
    go_log "log"

    "github.com/deatil/doak-fs/pkg/log"
    "github.com/deatil/doak-fs/pkg/log/handlers/json"
)

// 日志文件
var defaultLogFile = "./fs-log.log"

// 日志等级
var defaultLogLevel = "error"

var levelStrings = map[string]log.Level{
    "debug":   log.DebugLevel,
    "info":    log.InfoLevel,
    "warn":    log.WarnLevel,
    "warning": log.WarnLevel,
    "error":   log.ErrorLevel,
    "fatal":   log.FatalLevel,
}

type Fields = log.Fields

var logger *log.Entry
var once sync.Once

// 日志
// Logger().Debug(msg string)
func Logger() *log.Entry {
    once.Do(func() {
        logger = Manager(defaultLogFile, defaultLogLevel)
    })

    return logger
}

// 日志
// Debug | Info | Warn | Error | Fatal
// Debugf | Infof | Warnf | Errorf | Fatalf | Trace
func Manager(logPath string, level string) *log.Entry {
    logger := &log.Logger{}

    lf, err := os.OpenFile(logPath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0660)
    if err != nil {
        go_log.Fatalf("Failed to open log file: %v", err)
    }
    // defer lf.Close()

    logger.Handler = json.New(lf)
    logger.Level = GetLoggerLevel(level)

    loggerEntry := logger.WithFields(log.Fields{
        "type": "doak-fs",
    })

    return loggerEntry
}

// 获取日志等级
func GetLoggerLevel(name string) log.Level {
    if level, ok := levelStrings[name]; ok {
        return level
    }

    return log.ErrorLevel
}

// 设置日志等级
func SetLoggerLevel(name string) {
    if _, ok := levelStrings[name]; ok {
        defaultLogLevel = name
    }
}

// 设置日志文件
func SetLoggerFile(file string) {
    defaultLogFile = file
}
