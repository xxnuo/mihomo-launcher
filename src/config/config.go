package config

import (
	"bufio"
	_ "embed"
	"fmt"
	"github.com/xxnuo/mihomo-launcher/cfs"
	"github.com/xxnuo/mihomo-launcher/src/log"
	"github.com/xxnuo/mihomo-launcher/src/utils"
	"gopkg.in/yaml.v3"
	"os"
	"path/filepath"
	"reflect"
	"strings"
)

const (
	ConfigDirname          = ".config/mihomo" // 配置文件所在目录名
	ConfigFilename         = "config.yaml"    // 配置文件名
	LauncherConfigFilename = "launcher.yaml"  // 配置文件名
)

var (
	HomeDir                = "" // 用户目录: "C:\Users\username"
	ConfigDir              = "" // 用户目录: "C:\Users\username\.config\mihomo"
	ConfigFilePath         = "" // 配置文件路径: "C:\Users\username\.config\mihomo\config.yaml"
	LauncherConfigFilePath = "" // 配置文件路径: "C:\Users\username\.config\mihomo\config.yaml"
	Config                 config
	LauncherConfig         launcherConfig
	UIDir                  = "./ui" // 默认值，会被覆盖
	BinDir                 = ""
	UwpExePath             = ""
	CoreExePath            = ""
)

type config struct {
	ExternalController string `yaml:"external-controller"`
	ExternalUI         string `yaml:"external-ui"`
	Secret             string `yaml:"secret"`
}

type launcherConfig struct {
	EnableSystemProxy bool   `yaml:"enable_system_proxy"`
	Version           string `yaml:"version"`
}

// Init 初始化配置文件
func Init() {
	log.Infoln("初始化配置文件")

	// 初始化默认配置
	Config = config{
		ExternalController: "127.0.0.1:9090",
		Secret:             "mihomo",
		ExternalUI:         UIDir,
	}

	// 初始化配置文件
	ConfigDir = GetConfigDir()
	ConfigFilePath = filepath.Join(ConfigDir, ConfigFilename)
	log.Debugln("%+v", ConfigFilePath)
	// 判断配置文件是否存在
	if isExists := utils.Exists(ConfigFilePath); !isExists {
		if isEmpty := utils.IsDirEmpty(ConfigDir); isEmpty {
			// 配置文件不存在，目录为空释放所有文件
			err := cfs.ExtractFolder(ConfigDir, cfs.RootFolder) // 释放 cfs.FS 中 root 目录下的所有文件
			if err != nil {
				panic(err)
			}
		} else {
			// 配置文件不存在，目录不为空，释放默认配置文件
			err := cfs.ExtractFile(ConfigFilename, ConfigFilePath)
			if err != nil {
				panic(err)
			}
		}
	}

	LauncherConfigFilePath = filepath.Join(ConfigDir, LauncherConfigFilename)
	if isExists := utils.Exists(LauncherConfigFilePath); !isExists {
		// 配置文件不存在，目录不为空，释放默认配置文件
		err := cfs.ExtractFile(LauncherConfigFilename, LauncherConfigFilePath)
		if err != nil {
			panic(err)
		}
	}

	// 读取配置文件
	err := Read(ConfigFilePath, &Config)
	if err != nil {
		log.Errorln("Core 配置文件读取失败:%s", err)
	}

	err = LauncherRead()
	if err != nil {
		log.Errorln("Launcher 配置文件读取失败:%s", err)
	}

	// 写入配置
	err = Write(ConfigFilePath, &Config)
	if err != nil {
		log.Errorln("获取 Core 配置文件更新权限失败:%s", err)
	}

	err = LauncherWrite()
	if err != nil {
		log.Errorln("获取 Launcher 配置文件更新权限失败:%s", err)
	}

	// 确认 UI 文件存在
	UIDir = filepath.Join(ConfigDir, Config.ExternalUI, "index.html")
	if isExists := utils.Exists(UIDir); !isExists {
		err = cfs.ExtractFolder(filepath.Join(ConfigDir, Config.ExternalUI), cfs.UIFolder)
		if err != nil {
			panic(err)
		}
	}

	BinDir = filepath.Join(ConfigDir, cfs.BinFolderName)
	UwpExePath = filepath.Join(BinDir, cfs.UwpExeName)
	CoreExePath = filepath.Join(BinDir, cfs.CoreExeName)

	// 初始化完成
	log.Infoln("配置文件读取完成")

}

// GetConfigDir 获取配置文件路径
func GetConfigDir() string {
	HomeDir = utils.GetUserHomeDir()
	if HomeDir == "" {
		log.Errorln("无法获取用户目录")
	}
	defaultPath, _ := utils.EnsureDir(HomeDir, ConfigDirname)
	return defaultPath
}

// Read 读入配置文件
func Read(configFilePath string, readedConfig *config) error {
	configFileRaw, err := os.ReadFile(configFilePath)
	if err != nil {
		return err
	}
	err = yaml.Unmarshal(configFileRaw, &readedConfig)
	if err != nil {
		return err
	}
	return nil
}

// LauncherRead 读入配置文件
func LauncherRead() error {
	configFilePath, readedConfig := LauncherConfigFilePath, &LauncherConfig
	configFileRaw, err := os.ReadFile(configFilePath)
	if err != nil {
		return err
	}
	err = yaml.Unmarshal(configFileRaw, &readedConfig)
	if err != nil {
		return err
	}
	return nil
}

// Write 写入配置文件
func Write(configFilePath string, newConfig *config) error {
	newConfigMap := make(map[string]string)
	newConfigWriteFlag := make(map[string]bool)

	// 通过反射获取结构体中的yaml字段值和值内容
	vauleType := reflect.TypeOf(*newConfig)
	value := reflect.ValueOf(*newConfig)
	var k, v string
	for i := 0; i < vauleType.NumField(); i++ {
		k = vauleType.Field(i).Tag.Get("yaml")
		v = value.Field(i).String()
		newConfigMap[k] = v
		newConfigWriteFlag[k] = false
	}

	log.Debugln("%+v", newConfigMap)
	// map[external-controller:0.0.0.0:7899 external-ui:./ui/ secret:hi]

	// 打开文件以供读写
	file, err := os.OpenFile(configFilePath, os.O_RDWR, 0644)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return err
	}
	defer file.Close()

	// 创建一个用于写入文件的缓冲写入器
	writer := bufio.NewWriter(file)
	scanner := bufio.NewScanner(file)

	// 逐行读取文件内容
	var lines []string
	for scanner.Scan() {
		line := scanner.Text()

		// 检查行是否以特定字符串开头
		for k, v := range newConfigMap {
			if strings.HasPrefix(line, k) {
				// 如果匹配，则替换该行内容
				line = k + ": " + v
				newConfigWriteFlag[k] = true
			}
		}

		lines = append(lines, line)
	}

	// 清空文件内容
	err = file.Truncate(0)
	if err != nil {
		//fmt.Println("Error truncating file:", err)
		return err
	}

	// 将文件指针移动到文件开头
	_, err = file.Seek(0, 0)
	if err != nil {
		//fmt.Println("Error seek:", err)
		return err
	}

	// 将处理后的行写入文件
	for _, line := range lines {
		_, err := fmt.Fprintln(writer, line)
		if err != nil {
			//fmt.Println("Error writing to file:", err)
			return err
		}
	}

	// 将未写入的配置写入文件末尾
	for k, v := range newConfigWriteFlag {
		if !v {
			line := k + ": " + newConfigMap[k]
			_, err := fmt.Fprintln(writer, line)
			if err != nil {
				//fmt.Println("Error writing to file:", err)
				return err
			}
		}
	}

	// 刷新缓冲并确保所有写入操作完成
	if err := writer.Flush(); err != nil {
		//fmt.Println("Error flushing writer:", err)
		return err
	}

	if err := scanner.Err(); err != nil {
		//fmt.Println("Error scanning file:", err)
		return err
	}

	//fmt.Println("File updated successfully.")

	return nil
}

// LauncherWrite 写入配置文件
func LauncherWrite() error {
	configFilePath, newConfig := LauncherConfigFilePath, &LauncherConfig
	// 将结构体编码为 YAML
	data, err := yaml.Marshal(&newConfig)
	if err != nil {
		log.Errorln("error: %v", err)
	}

	// 将 YAML 数据写入到文件
	err = os.WriteFile(configFilePath, data, 0644)
	if err != nil {
		log.Errorln("error: %v", err)
	}

	//fmt.Println("File updated successfully.")

	return nil
}
