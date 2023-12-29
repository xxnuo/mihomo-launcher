//go:build darwin

package main

import _ "embed"

//go:embed winres/iconmac.icns
var iconData []byte

func IsAdmin() bool {
	return os.Geteuid() != 0
}

func RestartAsAdmin() {
	//fmt.Println("Restarting with admin privileges...")
	cmd := exec.Command("sudo", os.Args[0])
	cmd.Stdout = os.Stdout
	cmd.Stdin = os.Stdin
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		//fmt.Println("Error restarting:", err)
		return
	}
	os.Exit(0)
}
