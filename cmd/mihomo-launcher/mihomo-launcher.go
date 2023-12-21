package main

import (
	_ "embed"
	"github.com/xxnuo/clash-api/clash"
	"github.com/xxnuo/mihomo-launcher/internal/config"
	"github.com/xxnuo/mihomo-launcher/internal/log"
	"github.com/xxnuo/mihomo-launcher/internal/panel"
	"github.com/xxnuo/mihomo-launcher/internal/service"
	"github.com/xxnuo/mihomo-launcher/internal/utils"
	"github.com/xxnuo/systray"
	"time"
)

var RuntimeInfo struct {
	CoreVersion *clash.Version
	RawConfigs  clash.RawConfigs
	ClashPorts  clash.Ports

	ExtCtlPort int // 外部控制端口
}

func main() {
	err := error(nil)

	// 读取并检查配置文件
	err = config.LauncherInit()
	if err != nil {
		log.Errorln("Config: %w", err)
		panic(err)
	}

	// 检查服务状态
	err = service.EnsureServiceRunning()
	if err != nil {
		log.Errorln("Service: %w", err)
		panic(err)
	}

	// 启动内核并确保内核正常运行
	err = service.RunCore()
	if err != nil {
		log.Errorln("启动内核失败！请检查服务运行状态: %w", err)
		panic(err)
	}

	// 初始化 api
	clash.SetSecrete(config.LauncherConfig.CoreArgs.Secret)
	clash.SetURL("http://" + config.LauncherConfig.CoreArgs.ExtCtl)

	// 通过 api 获取信息
	connected := false
	for !connected {
		time.Sleep(3 * time.Second)
		log.Infoln("尝试与内核建立通信...")
		if utils.IsPortAvailable(config.ExtCtlPort) {
			// 外部控制端口未开启
			continue
		}
		RuntimeInfo.CoreVersion, err = clash.GetVersion()
		log.Debugln("%+v", RuntimeInfo.CoreVersion)

		if err == nil && RuntimeInfo.CoreVersion.Meta && RuntimeInfo.CoreVersion.Version != "" {
			connected = true
			break
		}
	}
	log.Infoln("与内核建立通信成功")

	// 根据配置初始化内核设置
	err = clash.SetTunEnable(config.LauncherConfig.EnableTun)
	if err != nil {
		log.Errorln("设置 Tun 失败: %w", err)
		config.LauncherConfig.EnableTun = false
	}

	// 获取内核配置
	RuntimeInfo.RawConfigs, err = clash.GetConfigs()
	RuntimeInfo.ClashPorts, err = clash.GetPorts(RuntimeInfo.RawConfigs)
	if err != nil {
		log.Errorln("获取内核配置失败: %w", err)
	}

	//log.Debugln("%+v", RuntimeInfo.RawConfigs)

	// 设置面板参数
	panel.Init(config.LauncherConfig.CoreArgs.ExtCtl, config.LauncherConfig.CoreArgs.Secret)

	// 添加托盘图标
	systray.Run(trayOnReady, trayOnExit, trayOnLClick, trayOnRClick)

	// 等待退出
	log.Infoln("Exit")
}
