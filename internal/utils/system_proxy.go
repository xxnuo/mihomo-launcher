package utils

import (
	"fmt"
	"github.com/wzshiming/sysproxy"
	"github.com/xxnuo/mihomo-launcher/internal/log"
)

// SetSystemProxy 设置系统代理
func SetSystemProxy(enable bool, port int) error {
	if enable {
		host := "127.0.0.1"
		url := fmt.Sprintf("%s:%d", host, port)
		log.Debugln("Set system proxy:%t %s", enable, url)
		return enableSystemProxy(url)
	} else {
		return disableSystemProxy()
	}
}

// enableSystemProxy 开启系统代理
func enableSystemProxy(url string) error {
	err := error(nil)
	noProxy := []string{"localhost", "127.*", "192.168.*", "10.*", "172.16.*", "<local>"}
	err = sysproxy.OnHTTP(url)
	err = sysproxy.OnHTTPS(url)
	err = sysproxy.OnNoProxy(noProxy)
	return err
}

// disableSystemProxy 关闭系统代理
func disableSystemProxy() error {
	err := error(nil)
	err = sysproxy.OffHTTP()
	err = sysproxy.OffHTTPS()
	err = sysproxy.OffNoProxy()
	return err
}
