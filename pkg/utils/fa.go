package utils

var faIcons = map[string]string{
    "folder":  "fa fa-folder-o",       //文件夹
    "file":    "fa fa-file-o",            //普通文件
    "image":   "fa fa-file-image-o",      //图片
    "xls":     "fa fa-file-excel-o",
    "word":    "fa fa-file-word-o",
    "ppt":     "fa fa-file-powerpoint-o",
    "pdf":     "fa fa-file-pdf-o",
    "code":    "fa fa-file-code-o",       //代码
    "archive": "fa fa-file-archive-o",    //压缩包
    "txt":     "fa fa-file-text-o",       //文本
    "audio":   "fa fa-file-audio-o",      //音频
    "video":   "fa fa-file-video-o",      //视频
    "apk":     "fa fa-android",           //安卓apk
    "exe":     "fa fa-beer",              //windows可执行文件
    "md":      "fa fa-file-text",
}

// 获取图标
func GetFaIcon(name string) string {
    if icon, ok := faIcons[name]; ok {
        return icon
    }

    return faIcons["file"]
}
