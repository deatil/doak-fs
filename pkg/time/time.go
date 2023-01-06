package time

import (
    "time"
    "github.com/jinzhu/now"

    "github.com/deatil/doak-fs/pkg/global"
)

var (
    DateFormat     = "2006-01-02"
    TimeFormat     = "15:04:05"
    DateTimeFormat = "2006-01-02 15:04:05"
)

/**
 * 时间格式化
 *
 * @create 2023-1-3
 * @author deatil
 */
type Time struct {
    time time.Time
}

// 设置时间
func (this Time) WithTime(t time.Time) Time {
    this.time = t

    return this
}

// 输出时间
func (this Time) ToTime() time.Time {
    return this.time
}

// 输出格式化
func (this Time) ToFormatString(format string) string {
    if this.time.IsZero() {
        return ""
    }

    return this.time.Format(format)
}

// 输出格式化时间
func (this Time) ToDateTimeString() string {
    if this.time.IsZero() {
        return ""
    }

    return this.time.Format(DateTimeFormat)
}

// 输出格式化时间
func (this Time) ToDateString() string {
    if this.time.IsZero() {
        return ""
    }

    return this.time.Format(DateFormat)
}

// 输出格式化时间
func (this Time) ToTimeString() string {
    if this.time.IsZero() {
        return ""
    }

    return this.time.Format(TimeFormat)
}

// 来源时间
func FromTime(t time.Time) Time {
    loc := timeLoc()

    return Time{now.New(t.In(loc)).Time}
}

// 来源时间
func FromTimestamp(timestamp int64) Time {
    return FromTime(time.Unix(timestamp, 0))
}

// 当前时间
func Now() Time {
    return FromTime(time.Now())
}

// 解析
func Parse(str string) Time {
    loc := timeLoc()

    date, _ := now.ParseInLocation(loc, str)

    return Time{date}
}

// 解析失败后抛出异常
func MustParse(str string) Time {
    loc := timeLoc()

    date := now.MustParseInLocation(loc, str)

    return Time{date}
}

// 时区
func timeLoc() *time.Location {
    timezone := global.Conf.App.TimeZone
    loc, _ := time.LoadLocation(timezone)

    return loc
}
