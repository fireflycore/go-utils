package network

import (
	"net"
	"strings"
)

// GetInternalNetworkIp 获取本机在局域网中的 IP 地址
// 原理：通过建立 UDP 伪连接（不发送数据）来自动选择合适的路由接口 IP
// 返回: 字符串形式的 IP 地址（如 "192.168.1.5"）。如果获取失败，返回 "127.0.0.1"
func GetInternalNetworkIp() string {
	// 通过建立 UDP “伪连接”获取本机路由选路后的本地地址，不会真正发送业务数据。
	dial, err := net.Dial("udp", "114.114.114.114:80")
	if err != nil {
		return "127.0.0.1"
	}
	defer dial.Close()
	addr := dial.LocalAddr().String()

	index := strings.LastIndex(addr, ":")
	return addr[:index]
}
