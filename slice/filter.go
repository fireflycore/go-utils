package array

// FilterSlice 过滤切片
// array: 待过滤的切片
// fn: 过滤函数，返回 true 保留，返回 false 丢弃
func FilterSlice[T any](array []T, fn func(index int, item T) bool) []T {
	var temp []T
	for index, item := range array {
		if fn(index, item) {
			temp = append(temp, item)
		}
	}
	return temp
}
