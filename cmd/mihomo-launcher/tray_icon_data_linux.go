//go:build linux

package main

import (
	_ "embed"
)

//go:embed winres/Meta_fix.png
var trayIconData []byte
