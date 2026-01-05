package structx

import (
	"reflect"
	"slices"
)

// DiffStruct 比较两个结构体的字段差异
// old: 旧结构体对象（或指针）
// new: 新结构体对象（或指针）
// ignore: 需要忽略比较的字段名列表
// 返回: 变更字段名到新值的映射 (map[string]interface{})
// 注意: 仅比较导出字段，且 old/new 必须是相同的结构体类型（或其指针）
func DiffStruct[O any, N any](old O, new N, ignore []string) map[string]interface{} {
	ro := reflect.ValueOf(old)
	rn := reflect.ValueOf(new)

	if ro.Kind() != reflect.Struct && (ro.Kind() != reflect.Ptr || ro.Elem().Kind() != reflect.Struct) {
		return nil
	}
	if rn.Kind() != reflect.Struct && (rn.Kind() != reflect.Ptr || rn.Elem().Kind() != reflect.Struct) {
		return nil
	}

	if ro.Kind() == reflect.Ptr {
		ro = ro.Elem()
	}
	if rn.Kind() == reflect.Ptr {
		rn = rn.Elem()
	}

	variant := make(map[string]interface{})

	for i := 0; i < rn.NumField(); i++ {
		k := rn.Type().Field(i).Name

		if slices.Contains(ignore, k) || !ro.FieldByName(k).IsValid() {
			continue
		}

		nv := rn.Field(i).Interface()
		ov := ro.FieldByName(k).Interface()

		if !reflect.DeepEqual(ov, nv) {
			variant[k] = nv
		}
	}

	return variant
}
