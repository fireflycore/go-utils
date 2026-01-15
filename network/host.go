package network

import (
	"errors"
	"net"
	"strconv"
	"strings"
)

// SplitHostPort 将 addr 解析为 host/port，addr 不带端口时返回 defaultPort。
func SplitHostPort(addr string, defaultPort string) (string, string, error) {
	// 空地址直接返回错误。
	if addr == "" {
		return "", "", errors.New("empty address")
	}

	// 优先尝试按 host:port 解析。
	host, port, err := net.SplitHostPort(addr)
	// 解析成功则做基本校验后返回。
	if err == nil {
		// host 不能为空。
		if host == "" {
			return "", "", errors.New("empty host")
		}
		// 端口为空时使用默认端口。
		if port == "" {
			return host, defaultPort, nil
		}
		// 校验端口为数字。
		if _, err := strconv.Atoi(port); err != nil {
			return "", "", err
		}
		// 返回解析出的 host 与 port。
		return host, port, nil
	}

	// 未带端口时，把整个 addr 当作 host。
	host = strings.TrimSpace(addr)
	// host 不能为空。
	if host == "" {
		return "", "", errors.New("empty host")
	}
	// 返回 host 与默认端口。
	return host, defaultPort, nil
}
