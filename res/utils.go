package res

import (
	"strings"
)

// Join embed.FS 的 Join 命令
//
// 经测试，embed.FS 使用 filepath.Join 命令在 Windows 会出现 Bug
// 因为 filepath.Join 是与平台相关的命令，而在 embed.FS 内部路径都是统一使用 "/" 分隔符
// 所以在 Windows 下使用 filepath.Join 会 读取 embed.FS 出现路径错误无法读取的情况
// 所以这里自定义一个 Join 命令，统一使用 "/" 分隔符
func Join(fsDirNames ...string) string {
	// 如果没有元素，返回空字符串
	if len(fsDirNames) == 0 {
		return ""
	}
	// 如果只有一个元素，直接返回该元素
	if len(fsDirNames) == 1 {
		return fsDirNames[0]
	}
	// 使用 strings.Join 将多个字符串连接起来
	return strings.Join(fsDirNames, Separator)
}

// AddRoot 为 embed.FS 路径添加根目录
// 如果已经有根目录，则不添加
//
// 例如：
//
//	path = "root/bin" -> "root/bin"
//	path = "bin" -> "root/bin"
//	path = "demo.yaml" -> "root/demo.yaml"
func AddRoot(fsPath string) string {
	if strings.HasPrefix(fsPath, RootFolder) {
		return fsPath
	} else {
		return Join(RootFolder, fsPath)
	}
}
