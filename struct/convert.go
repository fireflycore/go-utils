package object

import (
	"encoding/json"
	"fmt"
)

// StructConvert 结构体转换
// 将 source 结构体转换为 target 结构体，通过 JSON 序列化/反序列化实现
// 注意：这会忽略性能，仅用于非关键路径或简单转换
func StructConvert[S any, T any](source *S, target *T) {
	marshal, mErr := json.Marshal(source)
	if mErr != nil {
		fmt.Println(mErr.Error())
		return
	}
	if err := json.Unmarshal(marshal, &target); err != nil {
		fmt.Println(err.Error())
		return
	}
}
