package main

import (
	_ "embed"
	"fmt"
	"github.com/xxnuo/clash-api/clash"
	"github.com/xxnuo/mihomo-launcher/src/config"
	"github.com/xxnuo/mihomo-launcher/src/log"
	"github.com/xxnuo/mihomo-launcher/src/panel"
	"github.com/xxnuo/mihomo-launcher/src/utils"
	"github.com/xxnuo/systray"
	"time"
)

var RuntimeInfo struct {
	CoreVersion     *clash.Version
	RawConfigs      clash.RawConfigs
	TunEnabled      bool
	Ports           clash.Ports
	SystemProxyPort int
}

var Ports struct {
}

func main() {
	err := error(nil)
	//if !IsAdmin() {
	//	RestartAsAdmin()
	//}

	// 释放配置文件
	config.Init()

	// 检查配置文件

	// 检查端口冲突

	// 启动内核并确保内核正常运行
	// TODO

	// 初始化 api
	clash.SetSecrete(config.Config.Secret)

	// 通过 api 获取信息
	connected := false
	for !connected {
		RuntimeInfo.CoreVersion, err = clash.GetVersion()
		log.Debugln("%+v", RuntimeInfo.CoreVersion)

		if err == nil && RuntimeInfo.CoreVersion.Meta && RuntimeInfo.CoreVersion.Version != "" {
			connected = true
			break
		}
		log.Errorln("与内核通信失败，3s 后重新尝试")
		time.Sleep(3 * time.Second)

	}

	RuntimeInfo.RawConfigs, err = clash.GetConfigs()

	RuntimeInfo.TunEnabled, err = clash.IsTunEnabled(RuntimeInfo.RawConfigs)
	RuntimeInfo.Ports, err = clash.GetPorts(RuntimeInfo.RawConfigs)
	RuntimeInfo.SystemProxyPort = RuntimeInfo.Ports.MixedPort

	if err != nil {
		log.Errorln("获取内核配置失败: %s", err)
	}

	// 设置系统代理
	utils.SetSystemProxy(config.LauncherConfig.EnableSystemProxy, RuntimeInfo.SystemProxyPort)
	if err != nil {
		log.Errorln("Set system proxy error: %s", err)
	}

	// 设置面板参数
	panel.Init(config.Config.ExternalController, config.Config.Secret)
	// 添加托盘图标
	systray.Run(trayOnReady, trayOnExit, trayOnLClick, trayOnRClick)
	//println("exit")
}

func trayOnLClick(showMenu func() error) {
	panel.Open()
	// 不需要显示菜单了
	//_ = showMenu()
}

func trayOnRClick(showMenu func() error) { _ = showMenu() }

func trayOnReady() {

	systray.SetIcon(iconData)
	systray.SetTitle("mihomo-launcher")
	systray.SetTooltip(fmt.Sprintf("mihomo-launcher %s\n内核版本：%s", config.LauncherConfig.Version, RuntimeInfo.CoreVersion.Version))

	trayBtnOpenPanel := systray.AddMenuItem("打开面板", "打开控制面板")
	trayChkSetProxy := systray.AddMenuItemCheckbox("系统代理", "设置为系统代理", config.LauncherConfig.EnableSystemProxy)
	trayChkTunMode := systray.AddMenuItemCheckbox("Tun 模式", "开启 Tun 模式", RuntimeInfo.TunEnabled)
	systray.AddSeparator()
	trayBtnUWPLoopback := systray.AddMenuItem("UWP 程序网络修复", "修复 UWP 程序的网络连接")
	trayBtnRestart := systray.AddMenuItem("重启", "重启程序和内核")
	trayBtnQuit := systray.AddMenuItem("退出", "退出程序和内核")

	go func() {
		for {
			select {
			case <-trayBtnOpenPanel.ClickedCh:
				//	打开控制面板
				panel.Open()
			case <-trayChkSetProxy.ClickedCh:
				// 设置为系统代理
				if trayChkSetProxy.Checked() {
					trayChkSetProxy.Uncheck()
				} else {
					trayChkSetProxy.Check()
				}
				trayChkSetProxy.Disable()
				//code
				config.LauncherConfig.EnableSystemProxy = trayChkSetProxy.Checked()
				err := utils.SetSystemProxy(config.LauncherConfig.EnableSystemProxy, RuntimeInfo.SystemProxyPort)
				if err != nil {
					log.Errorln("Set system proxy error: %s", err)
				}

				trayChkSetProxy.Enable()

			case <-trayChkTunMode.ClickedCh:
				// 切换 Tun 模式
				if trayChkTunMode.Checked() {
					trayChkTunMode.Uncheck()
				} else {
					trayChkTunMode.Check()
				}
				trayChkTunMode.Disable()
				// code
				RuntimeInfo.TunEnabled = trayChkTunMode.Checked()
				err := clash.SetTunEnable(RuntimeInfo.TunEnabled)
				if err != nil {
					log.Errorln("Switch tun mode error: %s", err)
				}

				trayChkTunMode.Enable()

			case <-trayBtnUWPLoopback.ClickedCh:
				// 执行 enableLoopback.exe
				utils.OpenUWPLoopback(config.UwpExePath)

			case <-trayBtnRestart.ClickedCh:
				// 重启程序和内核
				_ = clash.Restart()
				// TODO 重启程序
				//systray.Quit()
			case <-trayBtnQuit.ClickedCh:
				// 退出程序和内核
				// TODO
				systray.Quit()
				return
			}
		}

	}()
}

func trayOnExit() {
	//println("exit")
	err := config.LauncherWrite()
	if err != nil {
		log.Errorln("获取 Launcher 配置文件更新权限失败: %s", err)
	}
}
