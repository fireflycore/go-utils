## Process (进程)

提供进程信号处理工具，用于实现优雅关闭。

### API

#### Watcher

- **签名**: `func Watcher(handle func())`
- **描述**: 阻塞监听系统信号（SIGTERM, SIGQUIT, SIGINT），在接收到信号时执行回调函数 `handle`。

### 示例

```go
import (
    "fmt"
    "github.com/fireflycore/go-utils/process"
)

func main() {
    // 启动服务...

    // 阻塞等待退出信号
    process.Watcher(func() {
        fmt.Println("Cleaning up resources...")
        // 执行清理逻辑，如关闭数据库连接
    })
    
    fmt.Println("Server exited")
}
```
