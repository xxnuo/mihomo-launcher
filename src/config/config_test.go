package config

import (
	"testing"
)

func TestGetConfigDir(t *testing.T) {
	t.Log(GetConfigDir())
}

func TestExtractFiles(t *testing.T) {
	Init()
	//err := extractFiles()
	//if err != nil {
	//	panic(err)
	//}
}

func TestInit(t *testing.T) {

	Init()
	//err := extractFiles()
	//if err != nil {
	//	panic(err)
	//}
}

func TestRead(t *testing.T) {
	newConfig := config{}
	Read("C:\\Users\\bigtear\\.config\\mihomo\\config.yaml", &newConfig)
	t.Logf("%+v", newConfig)
}

func TestWrite(t *testing.T) {
	newConfig := config{
		ExternalController: "0.0.0.0:7899",
		Secret:             "hi",
		ExternalUI:         "./ui/",
	}
	err := Write("C:\\Users\\bigtear\\.config\\mihomo\\config.yaml", &newConfig)
	if err != nil {
		t.Log(err)
	}
}
