package object

import (
	"reflect"
	"slices"
)

// DiffStruct 比较两个结构体，返回发生改变的字段
// old: 旧结构体
// new: 新结构体
// ignore: 忽略比较的字段名列表
// 返回: 变更字段名 -> 新值的映射
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
