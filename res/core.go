package res

import (
	"github.com/xxnuo/mihomo-launcher/internal/utils"
	"path/filepath"
	"runtime"
)

var (
	CoreExePath = "" // %USERPROFILE%\\.config\\mihomo\\bin\\mihomo.exe

	coreExeNameList = map[string]string{
		"windows": "mihomo.exe",
		"other":   "mihomo",
	}
)

func getCoreBinName() string {
	switch runtime.GOOS {
	case "windows":
		return coreExeNameList["windows"]
	default:
		return coreExeNameList["other"]
	}
}

// EnsureCoreExe 确保核心程序存在 == Init
// 成功后设置 CoreExePath
func EnsureCoreExe(targetDir string) error {
	err := error(nil)
	CoreExePath = filepath.Join(targetDir, getCoreBinName())

	if !utils.Exists(CoreExePath) {
		err = utils.WriteBytes(CoreExePath, coreExeBinary)
	}

	return err
}
