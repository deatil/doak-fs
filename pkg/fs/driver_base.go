package fs

/**
 * 驱动公共类
 *
 * @create 2023-2-17
 * @author deatil
 */
type DriverBase struct {}

// 是否为文件夹
func (this DriverBase) Basename(path string) string {
    return Filesystem.Basename(path)
}

// 是否为文件夹
func (this DriverBase) ParentPath(path string) string {
    if path == "" || path == "/" {
        return ""
    }

    parentPath := Filesystem.Dirname(path)
    parentPath = Filesystem.ToSlash(parentPath)

    return parentPath
}

func (this DriverBase) Extension(path string) string {
    return Filesystem.Extension(path)
}
