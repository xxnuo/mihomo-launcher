//go:build darwin

package main

import _ "embed"

//go:embed winres/iconmac.icns
var trayIconData []byte
