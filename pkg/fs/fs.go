package fs

import (
    "io"
)

// 接口
type IFs interface {
    Ls(directory string) []map[string]any
    LsDir(directory string) []map[string]any
    Read(path string) map[string]any
    Delete(paths ...string) error
    Exists(path string) bool
    IsFile(path string) bool
    IsDirectory(path string) bool
    Get(path string) (string, error)
    Put(path string, contents string) error
    CreateFile(path string) error
    CreateDir(path string) error
    Upload(rd io.Reader, path string, name string) error
    Rename(oldName string, newName string) error
    Move(oldName string, newName string) error
    Copy(oldName string, newName string) error

    Basename(path string) string
    ParentPath(path string) string
    Extension(path string) string
    FormatFile(path string) (string, error)
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

func (this Fs) Get(path string) (string, error) {
    return this.Driver.Get(path)
}

func (this Fs) Put(path string, contents string) error {
    return this.Driver.Put(path, contents)
}

func (this Fs) CreateFile(path string) error {
    return this.Driver.CreateFile(path)
}

func (this Fs) Basename(path string) string {
    return this.Driver.Basename(path)
}

func (this Fs) ParentPath(path string) string {
    return this.Driver.ParentPath(path)
}

func (this Fs) Extension(path string) string {
    return this.Driver.Extension(path)
}

func (this Fs) Upload(src io.Reader, path string, name string) error {
    return this.Driver.Upload(src, path, name)
}

func (this Fs) Rename(oldName string, newName string) error {
    return this.Driver.Rename(oldName, newName)
}

func (this Fs) Move(oldName string, newName string) error {
    return this.Driver.Move(oldName, newName)
}

func (this Fs) Copy(oldName string, newName string) error {
    return this.Driver.Copy(oldName, newName)
}

func (this Fs) FormatFile(path string) (string, error) {
    return this.Driver.FormatFile(path)
}

func (this Fs) CreateDir(path string) error {
    return this.Driver.CreateDir(path)
}
