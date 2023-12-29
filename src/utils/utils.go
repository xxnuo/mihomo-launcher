package utils

import (
	"errors"
	"fmt"
	"github.com/skratchdot/open-golang/open"
	"github.com/wzshiming/sysproxy"
	"github.com/xxnuo/mihomo-launcher/src/log"
	"net"
	"os"
	"path/filepath"
	"runtime"
)

func Switch(v bool) string {
	if v {
		return "开"
	} else {
		return "关"
	}
}

// GetParentDir 获取父目录
func GetParentDir(dir string) string {
	return filepath.Dir(dir)
}

// GetUserHomeDir 获取用户目录
func GetUserHomeDir() string {
	env := "HOME"
	if runtime.GOOS == "windows" {
		env = "USERPROFILE"
	} else if runtime.GOOS == "plan9" {
		env = "home"
	}
	return os.Getenv(env)
}

// EnsureDir 确保目录存在，如果不存在则创建，内置 Join 方法会自动拼接路径
// 不检查目录是否存在，直接创建，如存在直接返回
// 参数 elem 为可变参数，可以传入多个字符串自动拼接路径
func EnsureDir(elem ...string) (string, error) {
	path := filepath.Join(elem...)
	err := os.MkdirAll(path, os.ModePerm)
	if err != nil {
		return "", err
	}
	return path, nil
}

// IsDir checks if a given path is a directory.
func IsDir(path string) bool {
	fi, err := os.Stat(path)
	if err != nil {
		return false
	}
	return fi.IsDir()
}

// Exists 检查文件是否存在
func Exists(path string) bool {
	// 构建文件的完整路径
	fullPath, _ := filepath.Abs(path)
	// 使用os.Stat检查文件是否存在
	_, err := os.Stat(fullPath)
	if err == nil {
		// 文件存在
		return true
	}
	if os.IsNotExist(err) {
		// 文件不存在
		return false
	}
	// 其他错误
	return false
}

// IsDirEmpty 检查目录是否为空
func IsDirEmpty(path string) bool {
	dir, err := os.Open(path)
	if err != nil {
		return false
	}
	defer dir.Close()
	_, err = dir.Readdir(1) // 读取一个文件
	if err == nil {
		// 目录非空
		return false
	}
	if errors.Is(err, os.ErrNotExist) {
		// 目录不存在
		return false
	}
	// 目录为空
	return true
}

// OpenUWPLoopback 打开控制面板
func OpenUWPLoopback(exePath string) {
	err := open.Run(exePath)
	if err != nil {

		log.Errorln("Error opening enableLoopback.exe:", err)
	}
}

// SetSystemProxy 设置系统代理
func SetSystemProxy(enable bool, port int) error {
	host := "127.0.0.1"
	url := fmt.Sprintf("%s:%d", host, port)
	log.Debugln("Set system proxy:%t %s", enable, url)
	if enable {
		return enableSystemProxy(url)
	} else {
		return disableSystemProxy()
	}
}

// enableSystemProxy 开启系统代理
func enableSystemProxy(url string) error {
	err := error(nil)
	noProxy := []string{"localhost", "127.*", "192.168.*", "10.*", "172.16.*", "<local>"}
	err = sysproxy.OnHTTP(url)
	err = sysproxy.OnHTTPS(url)
	err = sysproxy.OnNoProxy(noProxy)
	return err
}

// disableSystemProxy 关闭系统代理
func disableSystemProxy() error {
	err := error(nil)
	err = sysproxy.OffHTTP()
	err = sysproxy.OffHTTPS()
	err = sysproxy.OffNoProxy()
	return err
}

// IsPortAvailable 检查端口是否可用
func IsPortAvailable(port int) bool {
	ln, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		return false
	}

	err = ln.Close()
	if err != nil {
		return false
	}
	return true
}
