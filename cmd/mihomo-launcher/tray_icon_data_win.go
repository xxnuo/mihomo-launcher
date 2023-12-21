//go:build windows

package main

import (
	_ "embed"
)

//go:embed winres/iconwin.ico
var trayIconData []byte
