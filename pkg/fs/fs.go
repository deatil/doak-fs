package fs

// 接口
type IFs interface {
    Ls(directory string) []map[string]any
    LsDir(directory string) []map[string]any
    Read(path string) map[string]any
    Delete(paths ...string) error
    Exists(path string) bool
    IsFile(path string) bool
    IsDirectory(path string) bool
    Get(path string, lock ...bool) (string, error)
    Put(path string, contents string, lock ...bool) error
    Chmod(path string, mode uint32) error
}

func New(driver IFs) Fs {
    return Fs{
        Driver: driver,
    }
}

/**
 * 文件管理器
 *
 * @create 2023-2-14
 * @author deatil
 */
type Fs struct {
    Driver IFs
}

func (this Fs) Ls(directory string) []map[string]any {
    return this.Driver.Ls(directory)
}

func (this Fs) LsDir(directory string) []map[string]any {
    return this.Driver.LsDir(directory)
}

func (this Fs) Read(path string) map[string]any {
    return this.Driver.Read(path)
}

func (this Fs) Delete(paths ...string) error {
    return this.Driver.Delete(paths...)
}

func (this Fs) Exists(path string) bool {
    return this.Driver.Exists(path)
}

func (this Fs) IsFile(path string) bool {
    return this.Driver.IsFile(path)
}

func (this Fs) IsDirectory(path string) bool {
    return this.Driver.IsDirectory(path)
}

func (this Fs) Get(path string, lock ...bool) (string, error) {
    return this.Driver.Get(path, lock...)
}

func (this Fs) Put(path string, contents string, lock ...bool) error {
    return this.Driver.Put(path, contents, lock...)
}

func (this Fs) Chmod(path string, mode uint32) error {
    return this.Driver.Chmod(path, mode)
}
