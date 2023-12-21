package panel

import (
	"fmt"
	"github.com/xxnuo/mihomo-launcher/internal/log"
	"github.com/xxnuo/mihomo-launcher/internal/utils"
)

// Open 打开控制面板
func Open() {
	url := fmt.Sprintf("http://%s/ui/#/Setup?hostname=%s&secret=%s", ExternalController, ExternalController, Secret)
	err := utils.OpenURL(url)
	if err != nil {
		log.Errorln("控制面板打开出错！:%s", err)
	}
}
