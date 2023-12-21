//go:build windows && (amd64 || 386)

package res

import (
	_ "embed"
)

//go:embed root/bin/enableLoopback.exe
var enableLoopbackBinary []byte
