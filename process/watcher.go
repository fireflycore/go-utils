package process

import (
	"os"
	"os/signal"
	"syscall"
)

// Watcher 监听系统信号，阻塞当前 goroutine 直到接收到终止信号。
// 支持的信号: SIGTERM, SIGQUIT, SIGINT
// handle: 接收到信号后执行的回调函数，通常用于资源清理
func Watcher(handle func()) {
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGINT)
	_ = <-ch

	if handle != nil {
		handle()
	}
}
