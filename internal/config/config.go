package config

import (
	"github.com/xxnuo/mihomo-launcher/internal/log"
	"github.com/xxnuo/mihomo-launcher/internal/utils"
)

var (
	HomeDir       = "" // %USERPROFILE%
	CoreConfigDir = "" // %USERPROFILE%\.config\mihomo
	UIDir         = "" // %USERPROFILE%\.config\mihomo\ui
	BinDir        = "" // %USERPROFILE%\.config\mihomo\bin"

	LauncherConfigFilePath = "./launcher.yaml" // %USERPROFILE%\.config\mihomo\launcher.yaml
	CoreConfigFilePath     = "./config.yaml"   // %USERPROFILE%\.config\mihomo\config.yaml

	LauncherConfig launcherConfig
	ExtCtlPort     int
)

// GetCoreConfigDir 获取配置文件路径
func GetCoreConfigDir() string {
	HomeDir = utils.GetUserHomeDir()
	if HomeDir == "" {
		log.Errorln("无法获取用户目录")
	}
	defaultPath, _ := utils.EnsureDir(HomeDir, ConfigDirname)
	return defaultPath
}
