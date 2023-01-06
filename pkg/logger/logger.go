package logger

import (
    "os"
    "sync"
    go_log "log"

    "github.com/deatil/doak-fs/pkg/log"
    "github.com/deatil/doak-fs/pkg/log/handlers/json"
)

// 日志文件
const defaultLogPath = "./fs-log.log"

type Fields = log.Fields

var logger *log.Entry
var once sync.Once

// 日志
// Logger().Debug(msg string)
func Logger() *log.Entry {
    once.Do(func() {
        logger = Manager(defaultLogPath)
    })

    return logger
}

// 日志
// Debug | Info | Warn | Error | Fatal
// Debugf | Infof | Warnf | Errorf | Fatalf | Trace
func Manager(logPath string) *log.Entry {
    logger := &log.Logger{}

    lf, err := os.OpenFile(logPath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0660)
    if err != nil {
        go_log.Fatalf("Failed to open log file: %v", err)
    }
    // defer lf.Close()

    logger.Handler = json.New(lf)
    logger.Level = log.WarnLevel

    loggerEntry := logger.WithFields(log.Fields{
        "type": "doak-fs",
    })

    return loggerEntry
}
