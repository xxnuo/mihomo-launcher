package res

import (
	"embed"
	"fmt"
	"github.com/xxnuo/mihomo-launcher/internal/utils"
	"io"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
)

//go:embed root/ui
var fsUI embed.FS

func EnsureUIFolder(targetDir string) error {
	_, err := utils.EnsureDir(targetDir)
	if err != nil {
		err = fmt.Errorf("获取 ui 目录失败:%w", err)
	}
	if isEmpty := utils.IsDirEmpty(targetDir); isEmpty {
		err = ExtractUIFolder(targetDir) // 释放内核文件
		if err != nil {
			err = fmt.Errorf("释放 ui 资源文件失败，无法正常使用控制面板:%w", err)
		}
	}
	return err
}

// ExtractUIFolder 写出 res.FS UI 目录中的所有文件到指定目录下
//
// 例如：
// root/ui/index.html -> targetDir/index.html
func ExtractUIFolder(targetDir string) error {
	embedFS := fsUI
	fsFolderPath := RootFolder + Separator + UIFolderName
	//fmt.Printf("0:src: %s\ntarget: %s\n", fsFolderPath, targetDir)
	err := fs.WalkDir(embedFS, fsFolderPath, func(path string, d fs.DirEntry, err error) error {
		//fmt.Printf("1:src: %s\ndir: %s\n", path, d.Name())
		if err != nil {
			return err
		}

		// 确认是指定目录下的文件
		if !strings.HasPrefix(path, fsFolderPath) {
			return nil
		}

		targetPath := filepath.Join(targetDir, strings.Replace(path, fsFolderPath, ".", 1))

		//fmt.Printf("2:src: %s\ntarget: %s\n", path, targetPath)
		if d.IsDir() {
			return os.MkdirAll(targetPath, d.Type())
		}

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
	})
	return err
}
