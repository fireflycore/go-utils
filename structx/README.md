# Structx (结构体扩展)

提供结构体转换与比较工具。

## API

### StructConvert

- **签名**: `func StructConvert[S any, T any](source *S, target *T)`
- **描述**: 通过 JSON 序列化/反序列化实现结构体之间的转换。
- **警告**: 性能较低，仅适用于非关键路径。建议高性能场景使用专门的转换库。

### DiffStruct

- **签名**: `func DiffStruct[O any, N any](old O, new N, ignore []string) map[string]interface{}`
- **描述**: 比较两个结构体，返回发生变更的字段名及其新值。支持忽略指定字段。

## 示例

```go
import (
    "fmt"
    "github.com/fireflycore/go-utils/structx"
)

type User struct {
    Name string
    Age  int
}

func main() {
    // StructConvert
    u1 := User{Name: "Alice", Age: 30}
    var u2 User
    structx.StructConvert(&u1, &u2)
    
    // DiffStruct
    oldUser := User{Name: "Bob", Age: 20}
    newUser := User{Name: "Bob", Age: 21}
    diff := structx.DiffStruct(oldUser, newUser, nil)
    fmt.Println(diff) // map[Age:21]
}
```
