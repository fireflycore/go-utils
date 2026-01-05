## Slices (切片)

提供泛型切片操作工具。

### API

#### DiffSlice

- **签名**: `func DiffSlice[T comparable](old, new []T) (add, remove []T)`
- **描述**: 比较两个切片的差异，返回新增和删除的元素列表。需要元素类型支持比较。

#### FilterSlice

- **签名**: `func FilterSlice[T any](array []T, fn func(index int, item T) bool) []T`
- **描述**: 根据过滤函数 `fn` 筛选切片中的元素，返回保留下来的元素组成的新切片。

### 示例

```go
import (
    "fmt"
    "github.com/fireflycore/go-utils/slices"
)

func main() {
    // DiffSlice 示例
    oldSlice := []int{1, 2, 3}
    newSlice := []int{2, 3, 4}
    add, remove := slices.DiffSlice(oldSlice, newSlice)
    fmt.Printf("Add: %v, Remove: %v\n", add, remove) // Add: [4], Remove: [1]

    // FilterSlice 示例
    nums := []int{1, 2, 3, 4, 5}
    evens := slices.FilterSlice(nums, func(i int, v int) bool {
        return v%2 == 0
    })
    fmt.Println("Evens:", evens) // Evens: [2 4]
}
```
