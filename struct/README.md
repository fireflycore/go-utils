## Struct (结构体)
提供结构体转换与比较工具。

- **StructConvert**
  - 签名: `func StructConvert[S any, T any](source *S, target *T)`
  - 描述: 通过 JSON 序列化/反序列化实现结构体之间的转换。注意：性能较低，仅适用于非关键路径。
- **DiffStruct**
  - 签名: `func DiffStruct[O any, N any](old O, new N, ignore []string) map[string]interface{}`
  - 描述: 比较两个结构体，返回发生变更的字段名及其新值。支持忽略指定字段。
