package utils

import (
	"fmt"
	"github.com/xxnuo/mihomo-launcher/internal/utils"
	"testing"
)

func TestIsPortAvailable(t *testing.T) {
	fmt.Println(utils.IsPortAvailable(7890))
	fmt.Println(utils.IsPortAvailable(33211))
	fmt.Println(utils.IsPortAvailable(33212))
	fmt.Println(utils.IsPortAvailable(33213))
}
