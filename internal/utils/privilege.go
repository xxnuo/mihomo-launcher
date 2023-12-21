package utils

import (
	"os"
	"runtime"
)

// IsAdmin 检查当前用户是否具有管理员权限
func IsAdmin() bool {
	switch runtime.GOOS {
	case "windows":
		return isAdminWindows()
	case "linux":
		return isAdminLinux()
	case "darwin":
		return isAdminMacOS()
	default:
		return false
	}
}

// isAdminWindows 检查当前用户是否具有管理员权限（Windows）
func isAdminWindows() bool {
	f, err := os.Open("\\\\.\\PHYSICALDRIVE0")
	defer f.Close()
	return err == nil
}

// isAdminLinux 检查当前用户是否具有管理员权限（Linux）
func isAdminLinux() bool {
	return os.Geteuid() == 0
}

// isAdminMacOS 检查当前用户是否具有管理员权限（macOS）
func isAdminMacOS() bool {
	return os.Geteuid() == 0
}
