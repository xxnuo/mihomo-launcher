package res

import (
	"github.com/xxnuo/mihomo-launcher/internal/log"
	"github.com/xxnuo/mihomo-launcher/internal/utils"
	"path/filepath"
)

// enableLoopbackExeName 这个只有一个 windows 386 架构
// 来自 https://github.com/Kuingsmile/uwp-tool
const enableLoopbackExeName = "enableLoopback.exe"

// ExecToolsUWP 执行 UWP 回环工具，仅支持 Windows x86 或 x64，在不支持的平台调用本函数不会有任何操作。
func ExecToolsUWP() {

	if enableLoopbackBinary == nil {
		// 不支持的平台
		return
	}

	targetDir, err := utils.GetTempDir()
	if err != nil {
		log.Errorln("Extract UWP tool failed!: %s", err)
	}

	targetExePath := filepath.Join(targetDir, enableLoopbackExeName)

	if err := utils.WriteBytes(targetExePath, enableLoopbackBinary); err != nil {
		log.Errorln("Extract UWP tool failed!: %s", err)
	}
	
	err = utils.RunExe(targetExePath)
	if err != nil {
		log.Errorln("Run UWP tool failed!: %s", err)
	}
	err = utils.DeleteDir(targetDir)
	if err != nil {
		return
	}

}
