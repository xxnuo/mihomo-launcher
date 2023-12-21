package utils

import (
	"github.com/skratchdot/open-golang/open"
)

// StartExe 不等待程序执行完成
func StartExe(exePath string) error { return open.Start(exePath) }

// RunExe 等待程序执行完成
func RunExe(exePath string) error { return open.Run(exePath) }

// OpenURL 打开网页
func OpenURL(url string) error { return open.Start(url) }
