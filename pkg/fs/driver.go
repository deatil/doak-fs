package fs

import(
    "sync"
)

func NewDriver() *Driver {
    return &Driver{
        drivers: make(map[string]DriverFunc),
    }
}

// 默认
var DefaultDriver = NewDriver()

type (
    DriverFunc = func() IFs
)

/**
 * 注册驱动
 *
 * @create 2023-2-15
 * @author deatil
 */
type Driver struct {
    // 锁定
    mu sync.RWMutex

    // 已注册数据
    drivers map[string]DriverFunc
}

// 添加
func (this *Driver) Add(name string, fs DriverFunc) {
    this.mu.Lock()
    defer this.mu.Unlock()

    if _, exists := this.drivers[name]; exists {
        delete(this.drivers, name)
    }

    this.drivers[name] = fs
}

func AddDriver(name string, fs DriverFunc) {
    DefaultDriver.Add(name, fs)
}

// 获取
func (this *Driver) Get(name string) IFs {
    this.mu.RLock()
    defer this.mu.RUnlock()

    fs, ok := this.drivers[name]
    if ok {
        return fs()
    }

    return nil
}

func GetDriver(name string) IFs {
    return DefaultDriver.Get(name)
}

// 判断
func (this *Driver) Has(name string) bool {
    this.mu.RLock()
    defer this.mu.RUnlock()

    _, exists := this.drivers[name]

    return exists
}

func HasDriver(name string) bool {
    return DefaultDriver.Has(name)
}

// 删除
func (this *Driver) Delete(name string) {
    this.mu.Lock()
    defer this.mu.Unlock()

    delete(this.drivers, name)
}

func DeleteDriver(name string) {
    DefaultDriver.Delete(name)
}

// 全部
func (this *Driver) All() map[string]DriverFunc {
    this.mu.Lock()
    defer this.mu.Unlock()

    return this.drivers
}

func AllDriver() {
    DefaultDriver.All()
}

// 全部驱动名称
func (this *Driver) Names() []string {
    this.mu.Lock()
    defer this.mu.Unlock()

    names := make([]string, 0)

    for name, _ := range this.drivers {
        names = append(names, name)
    }

    return names
}

func DriverNames() []string {
    return DefaultDriver.Names()
}
