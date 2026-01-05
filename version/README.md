# Version (版本)

提供简单的语义化版本管理工具。

## API

### Version

- 结构体: `Major`, `Minor`, `Patch`

### NewVersion

- **签名**: `func NewVersion(verStr string) (*Version, error)`
- **描述**: 解析版本字符串（支持带 'v' 前缀）。

### Methods

- **Increment**: `func (v *Version) Increment()`
  - 描述: 版本自增（逢 100 进位）。
- **Compare**: `func (v *Version) Compare(other *Version) int`
  - 描述: 版本比较（1: 大于, 0: 等于, -1: 小于）。
- **String**: `func (v *Version) String() string`
  - 描述: 格式化为 "vX.Y.Z" 形式。

## 示例

```go
import (
    "fmt"
    "github.com/fireflycore/go-utils/version"
)

func main() {
    v, _ := version.NewVersion("v1.0.99")
    v.Increment()
    fmt.Println(v.String()) // v1.1.0
    
    v2, _ := version.NewVersion("v2.0.0")
    fmt.Println(v.Compare(v2)) // -1
}
```
