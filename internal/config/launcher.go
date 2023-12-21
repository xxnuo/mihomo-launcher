package config

import (
	"fmt"
	"github.com/xxnuo/mihomo-launcher/internal/log"
	"github.com/xxnuo/mihomo-launcher/internal/utils"
	"github.com/xxnuo/mihomo-launcher/res"
	"gopkg.in/yaml.v3"
	"os"
	"path/filepath"
)

type launcherConfig struct {
	LogLevel          string   `yaml:"log-level"`
	EnableSystemProxy bool     `yaml:"enable-system-proxy"`
	EnableTun         bool     `yaml:"enable-tun"`
	GeoDataMode       bool     `yaml:"geo-data-mode"`
	CoreArgs          coreArgs `yaml:"core-args"`
}

type coreArgs struct {
	ExtCtl string `yaml:"ext-ctl"`
	ExtUI  string `yaml:"ext-ui"`
	Secret string `yaml:"secret"`
}

func (a *coreArgs) GetArgs() []string {
	return []string{
		"--ext-ctl", a.ExtCtl,
		"--ext-ui", a.ExtUI,
		"--secret", a.Secret,
	}
}

// LauncherInit 初始化配置文件
func LauncherInit() error {
	// 初始化默认配置
	LauncherConfig = launcherConfig{
		LogLevel:          log.DefaultLoglevel,
		EnableSystemProxy: false,
		EnableTun:         false,
		GeoDataMode:       false,
		CoreArgs: coreArgs{
			ExtCtl: DefaultExtCtl,
			ExtUI:  DefaultExtUI,
			Secret: DefaultSecret,
		},
	}
	// 初始化配置文件
	CoreConfigDir = GetCoreConfigDir()
	LauncherConfigFilePath = filepath.Join(CoreConfigDir, LauncherConfigFilename)
	BinDir = filepath.Join(CoreConfigDir, res.BinFolderName)
	UIDir = filepath.Join(CoreConfigDir, res.UIFolderName)

	// 判断配置文件是否存在
	_, err := utils.EnsureDir(CoreConfigDir)
	if err != nil {
		return fmt.Errorf("获取 %s 目录失败:%s", CoreConfigDir, err)

	}
	if isExists := utils.Exists(LauncherConfigFilePath); !isExists {
		// 配置文件不存在，目录不为空，释放默认配置文件
		err := LauncherWrite()
		if err != nil {
			return fmt.Errorf("默认配置写入失败:%s", err)
		}
	}
	// 配置文件存在，读取配置文件
	err = LauncherRead()
	if err != nil {
		return fmt.Errorf("配置文件读取失败:%w", err)
	}

	// 更新设置
	ExtCtlPort, err = utils.GetPort(LauncherConfig.CoreArgs.ExtCtl)
	if err != nil {
		return fmt.Errorf("内核 API 端口 %d 不是合法的 TCP 端口！: %w ", ExtCtlPort, err)
	}

	//UIDir = filepath.Join(CoreConfigDir, LauncherConfig.CoreArgs.ExtUI)

	log.SetLevelS(LauncherConfig.LogLevel)
	// DEBUG 调试用
	log.SetLevelS("debug")

	log.Debugln("UI 目录: %s", UIDir)

	// 确保内核可执行文件存在
	err = res.EnsureCoreExe(BinDir)
	if err != nil {
		return fmt.Errorf("释放内核失败:%w", err)
	}
	// 确保服务可执行文件存在
	err = res.EnsureServiceExes(BinDir)
	if err != nil {
		return fmt.Errorf("释放系统服务文件失败:%w", err)
	}
	// 确保 UI 文件存在
	err = res.EnsureUIFolder(UIDir)
	if err != nil {
		log.Errorln("释放 UI 资源失败，无法使用控制面板: %w", err)
	}
	// 初始化完成
	log.Infoln("启动器初始化完成")
	return nil
}

// LauncherWrite 写入配置文件
func LauncherWrite() error {
	configFilePath, newConfig := LauncherConfigFilePath, &LauncherConfig
	// 将结构体编码为 YAML
	data, err := yaml.Marshal(&newConfig)
	if err != nil {
		return err
	}
	// 将 YAML 数据写入到文件
	err = os.WriteFile(configFilePath, data, 0644)
	if err != nil {
		return err
	}
	return nil
}

// LauncherRead 读入配置文件
func LauncherRead() error {
	configFilePath, readConfig := LauncherConfigFilePath, &LauncherConfig
	configFileRaw, err := os.ReadFile(configFilePath)
	if err != nil {
		return err
	}
	err = yaml.Unmarshal(configFileRaw, &readConfig)
	if err != nil {
		return err
	}
	LauncherConfig.CoreArgs.ExtUI = filepath.Join(CoreConfigDir, LauncherConfig.CoreArgs.ExtUI)
	return nil
}
