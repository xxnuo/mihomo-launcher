package config

// 读取内核配置文件，暂未使用

import (
	"gopkg.in/yaml.v3"
	"os"
)

// CoreRead 读入配置文件
func CoreRead(configFilePath string, readConfig *coreConfig) error {
	configFileRaw, err := os.ReadFile(configFilePath)
	if err != nil {
		return err
	}
	err = yaml.Unmarshal(configFileRaw, &readConfig)
	if err != nil {
		return err
	}
	return nil
}
