package panel

import (
	"fmt"
	"github.com/skratchdot/open-golang/open"
)

var (
	ExternalController = "127.0.0.1"
	Secret             = ""
)

// Init 初始化控制面板参数
func Init(externalController string, secret string) {
	ExternalController = externalController
	Secret = secret
}

// Open 打开控制面板
func Open() {
	url := fmt.Sprintf("http://%s/ui/#/Setup?hostname=%s&secret=%s", ExternalController, ExternalController, Secret)
	err := open.Run(url)
	if err != nil {
		fmt.Println("Error opening webpage:", err)
	}
}
