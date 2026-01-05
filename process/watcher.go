package process

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

// Watcher 监听系统信号，阻塞当前 goroutine 直到接收到终止信号
// 支持的信号: SIGTERM, SIGQUIT, SIGINT
// handle: 接收到信号后执行的回调函数，通常用于资源清理
func Watcher(handle func()) {
	fmt.Println("----- watch process start -----")

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGINT)
	// 接收信号
	_ = <-ch

	if handle != nil {
		handle()
	}

	fmt.Println("----- watch process end -----")
}
