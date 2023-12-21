package utils

import (
	"errors"
	"os"
	"path/filepath"
	"runtime"
)

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

// GetTempDir 创建临时文件夹
func GetTempDir() (string, error) {
	tempDir, err := os.MkdirTemp("", "mihomo-*")
	if err != nil {
		return "", err
	}
	return tempDir, nil
}

// DeleteDir 删除文件夹
func DeleteDir(path string) error {
	return os.RemoveAll(path)
}
