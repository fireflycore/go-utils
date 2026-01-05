## File (文件)

提供简单的文件读写辅助功能。

### API

#### ReadRemoteFile

- **签名**: `func ReadRemoteFile(url string) ([]byte, error)`
- **描述**: 读取远程 URL 的文件内容。
- **注意**: 全量加载到内存，不适合大文件。

#### WriteLocalFile

- **签名**: `func WriteLocalFile(path string, bytes []byte) error`
- **描述**: 将字节内容写入本地文件。
- **特性**: 自动创建缺失的目录，覆盖已存在的文件。

### 示例

```go
import (
    "fmt"
    "github.com/fireflycore/go-utils/file"
)

func main() {
    // 读取远程
    data, err := file.ReadRemoteFile("https://example.com/data.json")
    if err != nil {
        panic(err)
    }

    // 写入本地
    err = file.WriteLocalFile("./data/backup.json", data)
    if err != nil {
        panic(err)
    }
}
```
