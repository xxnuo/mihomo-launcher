package tests

import (
	"fmt"
	"github.com/xxnuo/mihomo-launcher/internal/service"
	"testing"
)

func TestCheck(t *testing.T) {
	rsp, err := service.CheckService()
	fmt.Printf("%+v %+v\n", rsp, err)
}
