package utils

import (
	"os"
	"path/filepath"
)

// WriteBytes 将 bytes 写入到 targetFilePath
func WriteBytes(targetFilePath string, bytes []byte) error {
	_, err := EnsureDir(filepath.Dir(targetFilePath))
	if err != nil {
		return err
	}
	return os.WriteFile(targetFilePath, bytes, os.ModePerm)
}
