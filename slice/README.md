## Slice (切片)
提供泛型切片操作工具。

- **DiffSlice**
  - 签名: `func DiffSlice[T comparable](old, new []T) (add, remove []T)`
  - 描述: 比较两个切片的差异，返回新增和删除的元素列表。
- **FilterSlice**
  - 签名: `func FilterSlice[T any](array []T, fn func(index int, item T) bool) []T`
  - 描述: 根据过滤函数 `fn` 筛选切片中的元素，返回保留下来的元素组成的新切片。
