package object

import (
	"encoding/json"
	"fmt"
)

// StructConvert 结构体转换工具
// source: 源结构体指针
// target: 目标结构体指针
// 原理：通过 JSON 序列化 source 再反序列化到 target
// 警告：性能较差，仅适用于非关键路径或对性能不敏感的场景
// 建议：高性能场景请使用手动赋值或专门的转换库（如 goverter）
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
