# Network (网络)

提供网络相关的实用函数。

## API

### GetInternalNetworkIp

- **签名**: `func GetInternalNetworkIp() string`
- **描述**: 获取本机在局域网中的 IP 地址。通过建立 UDP 伪连接（不发送数据）来自动选择合适的路由接口 IP。

## 示例

```go
import (
    "fmt"
    "github.com/fireflycore/go-utils/network"
)

func main() {
    ip := network.GetInternalNetworkIp()
    fmt.Println("Local IP:", ip)
}
```
