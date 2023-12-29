//go:build windows

package main

import (
	_ "embed"
	"os"
	"os/exec"
	"syscall"
)

//go:embed winres/iconwin.ico
var iconData []byte

func IsAdmin() bool {
	_, err := os.Open("\\\\.\\PHYSICALDRIVE0")
	return err == nil
}

func RestartAsAdmin() {
	//fmt.Println("Restarting as admin...")
	cmd := exec.Command("runas", "/user:Administrator", os.Args[0])
	cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
	if err := cmd.Start(); err != nil {
		//fmt.Println("Error restarting:", err)
		return
	}
	os.Exit(0)
}
