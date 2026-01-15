# Network (网络)

提供网络相关的实用函数。

## API

### GetInternalNetworkIp

- **签名**: `func GetInternalNetworkIp() string`
- **描述**: 获取本机在局域网中的 IP 地址。通过建立 UDP 伪连接（不发送数据）来自动选择合适的路由接口 IP。

### SplitHostPort

- **签名**: `func SplitHostPort(addr string, defaultPort string) (host string, port string, err error)`
- **描述**: 将 `addr` 解析为 `host/port`。当 `addr` 不带端口时返回 `defaultPort` 作为兜底端口。
- **说明**:
  - `defaultPort` 仅在 `addr` 未显式包含端口时生效；当 `addr` 为 `host:port` 时，返回的 `port` 以 `addr` 为准。
  - 该方法不会修改你的原始地址字符串，只是返回解析结果，便于你在不同场景分别使用：
    - 建立连接通常需要 `host` 与 `port`
    - TLS 的 `ServerName` 只需要 `host`（不能包含端口）
- **错误**:
  - `addr` 为空：返回 `empty address`
  - `host` 为空（例如 `:5432`）：返回 `empty host`
  - `port` 非数字：返回 `strconv.Atoi` 的解析错误
- **注意**:
  - IPv6 使用 `net.SplitHostPort` 规则：带端口时必须写成 `[::1]:5432`

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

### SplitHostPort 示例

```go
import (
    "fmt"
    "github.com/fireflycore/go-utils/network"
)

func main() {
    host, port, _ := network.SplitHostPort("db.example.com:5432", "5432")
    fmt.Println(host, port) // db.example.com 5432

    host, port, _ = network.SplitHostPort("db.example.com", "5432")
    fmt.Println(host, port) // db.example.com 5432

    host, port, _ = network.SplitHostPort("[::1]:5432", "5432")
    fmt.Println(host, port) // ::1 5432
}
```
