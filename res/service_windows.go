//go:build windows && (amd64 || 386)

package res

import (
	_ "embed"
)

var (
	//go:embed root/bin/install-service.exe
	installServiceBinary []byte

	//go:embed root/bin/uninstall-service.exe
	uninstallServiceBinary []byte

	//go:embed root/bin/clash-verge-service.exe
	serviceBinary []byte
)
