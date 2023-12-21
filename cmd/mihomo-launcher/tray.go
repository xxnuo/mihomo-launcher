package main

import (
	"fmt"
	"github.com/xxnuo/clash-api/clash"
	"github.com/xxnuo/mihomo-launcher/internal/config"
	"github.com/xxnuo/mihomo-launcher/internal/log"
	"github.com/xxnuo/mihomo-launcher/internal/panel"
	"github.com/xxnuo/mihomo-launcher/internal/service"
	"github.com/xxnuo/mihomo-launcher/internal/utils"
	"github.com/xxnuo/mihomo-launcher/res"
	"github.com/xxnuo/systray"
	"runtime"
)

func trayOnLClick(showMenu func() error) { panel.Open() }

func trayOnRClick(showMenu func() error) { _ = showMenu() }

func trayOnExit() {
	//println("exit")
	err := config.LauncherWrite()
	if err != nil {
		log.Errorln("获取 Launcher 配置文件更新权限失败: %s", err)
	}
	err = utils.SetSystemProxy(false, 0)
	if err != nil {
		log.Errorln("Set system proxy error: %s", err)
	}
	systray.Quit()
}

func trayOnReady() {
	err := error(nil)
	systray.SetIcon(trayIconData)
	systray.SetTitle("mihomo-launcher")
	systray.SetTooltip(fmt.Sprintf("mihomo-launcher: %s\n内核: %s", Version, RuntimeInfo.CoreVersion.Version))

	mbOpenPanel := systray.AddMenuItem("打开面板", "打开控制面板")
	mcSetProxy := systray.AddMenuItemCheckbox("系统代理", "设置为系统代理", config.LauncherConfig.EnableSystemProxy)

	//设置系统代理
	err = utils.SetSystemProxy(config.LauncherConfig.EnableSystemProxy, RuntimeInfo.ClashPorts.MixedPort)
	if err != nil {
		log.Errorln("设置系统代理失败: %w", err)
	}

	mcTunMode := systray.AddMenuItemCheckbox("Tun 模式", "开启 Tun 模式", config.LauncherConfig.EnableTun)

	systray.AddSeparator()

	msTools := systray.AddMenuItem("工具", "工具")
	mbToolUWP := msTools.AddSubMenuItem("UWP 程序网络修复", "修复 UWP 程序的网络连接")

	if runtime.GOOS != "windows" {
		msTools.Disable()
	}

	msCore := systray.AddMenuItem("内核", "内核")
	mbCoreStart := msCore.AddSubMenuItem("启动", "启动")
	mbCoreStop := msCore.AddSubMenuItem("停止", "停止")
	mbCoreRestart := msCore.AddSubMenuItem("重启", "重启")

	if service.CoreStatus == service.Running {
		mbCoreStart.Disable()
	} else {
		mbCoreStop.Disable()
	}

	mbQuitAll := systray.AddMenuItem("关闭程序和内核", "退出程序和内核")
	mbQuitLauncher := systray.AddMenuItem("退出程序", "仅退出程序")

	go func() {
		for {
			select {
			case <-mbOpenPanel.ClickedCh:
				//	打开控制面板
				panel.Open()
			case <-mcSetProxy.ClickedCh:
				// 设置为系统代理
				mcSetProxy.Disable()

				if mcSetProxy.Checked() {
					mcSetProxy.Uncheck()
				} else {
					mcSetProxy.Check()
				}
				//code
				config.LauncherConfig.EnableSystemProxy = mcSetProxy.Checked()
				err := utils.SetSystemProxy(config.LauncherConfig.EnableSystemProxy, RuntimeInfo.ClashPorts.MixedPort)
				if err != nil {
					log.Errorln("Set system proxy error: %s", err)
				}

				mcSetProxy.Enable()

			case <-mcTunMode.ClickedCh:
				// 切换 Tun 模式
				mcTunMode.Disable()

				if mcTunMode.Checked() {
					mcTunMode.Uncheck()
				} else {
					mcTunMode.Check()
				}
				// code
				config.LauncherConfig.EnableTun = mcTunMode.Checked()
				err := clash.SetTunEnable(config.LauncherConfig.EnableTun)
				if err != nil {
					log.Errorln("Switch tun mode error: %s", err)
				}

				mcTunMode.Enable()

			case <-mbToolUWP.ClickedCh:
				// 执行 enableLoopback.exe
				res.ExecToolsUWP()

			case <-mbCoreStart.ClickedCh:
				// 手动启动内核
				err := service.RunCore()
				if err != nil {
					mbCoreStart.Enable()
					mbCoreStop.Disable()
					log.Errorln("Start core error: %s", err)
				} else {
					mbCoreStart.Disable()
					mbCoreStop.Enable()
				}
			case <-mbCoreStop.ClickedCh:
				// 手动停止内核
				err := service.StopCore()
				if err != nil {
					mbCoreStart.Enable()
					mbCoreStop.Disable()
					log.Errorln("Stop core error: %s", err)
				} else {
					mbCoreStart.Enable()
					mbCoreStop.Disable()
				}
			case <-mbCoreRestart.ClickedCh:
				// 手动重启内核
				err := service.RestartCore()
				if err != nil {
					log.Errorln("Restart core error: %s", err)
				} else {
					mbCoreStart.Disable()
					mbCoreStop.Enable()
				}
			case <-mbQuitAll.ClickedCh:
				// 关闭程序和内核
				err := service.StopCore()
				if err != nil {
					log.Errorln("Stop core error: %s", err)
				}
				trayOnExit()
			case <-mbQuitLauncher.ClickedCh:
				// 关闭程序
				trayOnExit()
			}
			err := config.LauncherWrite()
			if err != nil {
				log.Debugln("配置文件更新失败: %s", err)
			} else {
				log.Debugln("配置文件更新成功")
			}
		}

	}()
}
