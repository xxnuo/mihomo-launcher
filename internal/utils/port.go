package utils

import (
	"bufio"
	"fmt"
	"net"
	"net/http"
	"time"
)

// IsPortAvailable 检查端口是否未被占用
func IsPortAvailable(port int) bool {
	conn, err := net.DialTimeout("tcp", fmt.Sprintf("localhost:%d", port), time.Second)
	if err != nil {
		return true
	}
	defer conn.Close()
	// 发送一个HTTP请求到连接的端口
	_, err = fmt.Fprintf(conn, "GET / HTTP/1.0\r\n\r\n")
	if err != nil {
		return true
	}
	// 读取响应
	resp, err := http.ReadResponse(bufio.NewReader(conn), nil)
	if err != nil {
		return true
	}
	defer resp.Body.Close()
	// 判断HTTP响应状态码
	if resp.StatusCode == http.StatusOK {
		return true
	}
	return false
}

// GetPort 从 IP:PORT 文本中取得合法端口号
func GetPort(addr string) (int, error) {
	_, port, err := net.SplitHostPort(addr)
	if err != nil {
		return 0, err
	}
	return net.LookupPort("tcp", port)
}
