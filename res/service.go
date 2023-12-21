package res

import (
	"fmt"
	"github.com/xxnuo/mihomo-launcher/internal/utils"
	"path/filepath"
)

var (
	InstallServiceExePath   = "" // %USERPROFILE%\\.config\\mihomo\\bin\\install-service.exe
	UninstallServiceExePath = "" // %USERPROFILE%\\.config\\mihomo\\bin\\uninstall-service.exe
	ServiceExePath          = "" // %USERPROFILE%\\.config\\mihomo\\bin\\clash-verge-service.exe
)

// EnsureServiceExes 确保服务程序存在 == Init
// 成功后设置 InstallServiceExePath, UninstallServiceExePath, ServiceExePath
func EnsureServiceExes(targetDir string) error {
	if serviceBinary == nil {
		// 不支持的平台
		return fmt.Errorf("unsupported platform")
	}

	err := error(nil)
	InstallServiceExePath = filepath.Join(targetDir, InstallServiceExeName)
	UninstallServiceExePath = filepath.Join(targetDir, UninstallServiceExeName)
	ServiceExePath = filepath.Join(targetDir, ServiceExeName)

	if !utils.Exists(InstallServiceExePath) {
		e := utils.WriteBytes(InstallServiceExePath, installServiceBinary)
		if e != nil {
			err = e
		}
	}
	if !utils.Exists(UninstallServiceExePath) {
		e := utils.WriteBytes(UninstallServiceExePath, uninstallServiceBinary)
		if e != nil {
			err = e
		}
	}
	if !utils.Exists(ServiceExePath) {
		e := utils.WriteBytes(ServiceExePath, serviceBinary)
		if e != nil {
			err = e
		}
	}
	return err
}

func ExecInstallService() error { return utils.RunExe(InstallServiceExePath) }
func ExecUninstallService() error {
	return utils.RunExe(UninstallServiceExePath)
}
