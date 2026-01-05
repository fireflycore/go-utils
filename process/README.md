## Process (进程)
提供进程信号处理工具。

- **Watcher**
  - 签名: `func Watcher(handle func())`
  - 描述: 阻塞监听系统信号（SIGTERM, SIGQUIT, SIGINT），在接收到信号时执行回调函数 `handle`，常用于优雅关闭服务。