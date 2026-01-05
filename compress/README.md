# Compress (压缩)

提供数据压缩与解压的通用接口及实现。

## 接口定义

```go
type Compress interface {
    Compress(data []byte) ([]byte, error)
    Decompress(data []byte) ([]byte, error)
}
```

## 支持的算法

### GZIP

使用 Go 标准库 `compress/gzip` 实现。

- **NewGZIP**
  - 签名: `func NewGZIP() Compress`
  - 描述: 创建 GZIP 压缩器实例。

## 示例

```go
import "github.com/fireflycore/go-utils/compress"

func main() {
    c := compress.NewGZIP()
    
    // 压缩
    compressed, err := c.Compress([]byte("hello world"))
    if err != nil {
        panic(err)
    }

    // 解压
    original, err := c.Decompress(compressed)
    if err != nil {
        panic(err)
    }
}
```
