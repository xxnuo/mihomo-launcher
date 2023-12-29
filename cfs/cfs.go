package cfs

import (
	"embed"
	"github.com/xxnuo/mihomo-launcher/src/log"
	"github.com/xxnuo/mihomo-launcher/src/utils"
	"io"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
)

const (
	Separator     = `/`
	RootFolder    = "root"
	UIFolderName  = "ui"
	BinFolderName = "bin"
	UwpExeName    = "enableLoopback.exe"
	CoreExeName   = "mihomo-windows-amd64.exe"

	UIFolder    = RootFolder + Separator + UIFolderName
	BinFolder   = RootFolder + Separator + BinFolderName
	UwpExePath  = RootFolder + Separator + BinFolder + Separator + UwpExeName
	CoreExePath = RootFolder + Separator + BinFolder + Separator + CoreExeName
)

// FS 结构:
//
//	bin/
//	ui/
//
//go:embed root
var FS embed.FS

func trimFolder(path string, folderPath string) string {
	if strings.HasPrefix(path, folderPath) {
		return strings.Replace(path, folderPath, ".", 1)
	} else {
		return path
	}
}

func addRoot(path string) string {
	if strings.HasPrefix(path, RootFolder) {
		return path
	} else {
		return Join(RootFolder, path)
	}
}

/*
	Join embed.FS 的 Join 命令

经测试，embed.FS 使用 filepath.Join 命令会出现 Bug
因为 filepath.Join 是与平台相关的命令，而在 embed.FS 内部路径都是统一使用 "/" 分隔符
所以在 Windows 下使用 filepath.Join 会 读取 embed.FS 出现路径错误无法读取的情况
例如 Windows 下：

fileName1 := filepath.Join("dir", "a.txt") => "dir\\a.txt"
embedFS.Open(fileName1) => 找不到文件
fileName2 := "dir/a.txt"
embedFS.Open(fileName2) => ok

所以这里自定义一个 Join 命令，统一使用 "/" 分隔符
*/
func Join(elem ...string) string {
	// 如果没有元素，返回空字符串
	if len(elem) == 0 {
		return ""
	}
	// 如果只有一个元素，直接返回该元素
	if len(elem) == 1 {
		return elem[0]
	}
	// 使用 strings.Join 将多个字符串连接起来
	return strings.Join(elem, Separator)
}

// ReadFile 读取 cfs.FS 中的单个文件
func ReadFile(filePath string) ([]byte, error) {
	embedFS := FS
	file, err := embedFS.Open(addRoot(filePath))
	if err != nil {
		return nil, err
	}
	defer file.Close()
	return io.ReadAll(file)
}

// ExtractFolder 写出 cfs.FS root/ 目录中的所有文件
func ExtractFolder(targetDir string, folderName string) error {
	embedFS := FS
	srcFS, err := fs.Sub(embedFS, ".")
	if err != nil {
		return err
	}

	err = fs.WalkDir(srcFS, ".", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		// 确认是指定目录下的文件
		if strings.HasPrefix(path, folderName) {
			targetPath := filepath.Join(targetDir, trimFolder(path, folderName))
			if d.IsDir() {
				return os.MkdirAll(targetPath, d.Type())
			} else {
				srcFile, err := embedFS.Open(path)
				if err != nil {
					return err
				}
				defer srcFile.Close()

				targetFile, err := os.Create(targetPath)
				if err != nil {
					return err
				}
				defer targetFile.Close()

				_, err = io.Copy(targetFile, srcFile)
				return err
			}
		}
		return nil
	})

	return err
}

// ExtractFile 写出 cfs.FS 中的单个文件
func ExtractFile(srcFilePath string, targetFilePath string) error {
	log.Infoln("%s %s", srcFilePath, targetFilePath)
	embedFS := FS
	srcFilePath = addRoot(srcFilePath)

	srcFile, err := embedFS.Open(srcFilePath)
	if err != nil {
		return err
	}
	log.Infoln("源文件读取成功")

	utils.EnsureDir(filepath.Dir(targetFilePath))
	targetFile, _ := os.Create(targetFilePath)

	log.Infoln("模板文件创建成功")
	defer targetFile.Close()

	_, err = io.Copy(targetFile, srcFile)
	return err
}
