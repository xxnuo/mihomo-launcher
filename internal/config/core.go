package config

// 未使用，目前通过启动内核的命令行传入参数覆盖配置内容

import (
	"github.com/xxnuo/mihomo-launcher/internal/log"
	"github.com/xxnuo/mihomo-launcher/internal/utils"
	"path/filepath"
)

type coreConfig struct {
	ExternalController string `yaml:"external-controller"`
	ExternalUI         string `yaml:"external-ui"`
	Secret             string `yaml:"secret"`
}

var (
	CoreConfig coreConfig
)

// coreInit 使用读取配置文件方式初始化配置文件
// 需要在 LauncherInit 初始化配置完成后执行
func coreInit() {
	log.Infoln("初始化内核配置文件")

	// 初始化默认配置
	CoreConfig = coreConfig{
		ExternalController: "0.0.0.0:9090",
		ExternalUI:         "./ui",
		Secret:             "mihomo-launcher",
	}

	// 初始化配置文件
	CoreConfigDir = GetCoreConfigDir()
	CoreConfigFilePath = filepath.Join(CoreConfigDir, CoreConfigFilename)

	// 判断配置文件是否存在
	if isExists := utils.Exists(LauncherConfigFilePath); !isExists {
		// 配置文件不存在，目录不为空，释放默认配置文件
		err := CoreWrite(CoreConfigFilePath, &CoreConfig)
		if err != nil {
			log.Errorln("默认配置写入失败:%s", err)
			panic(err)
		}
	}

	// 读取 Core 配置文件
	log.Debugln("Read config: %+v", CoreConfigFilePath)
	err := CoreRead(CoreConfigFilePath, &CoreConfig)
	if err != nil {
		log.Errorln("Launcher 配置文件读取失败:%s", err)
	}

	// 初始化完成
	log.Infoln("内核配置文件读取完成")
}
