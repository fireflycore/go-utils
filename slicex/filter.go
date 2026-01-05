package slicex

// FilterSlice 过滤切片中的元素
// array: 待过滤的源切片
// fn: 过滤函数，接收索引和元素值，返回 true 表示保留，false 表示丢弃
// 返回: 包含所有保留元素的新切片
func FilterSlice[T any](array []T, fn func(index int, item T) bool) []T {
	var temp []T
	for index, item := range array {
		if fn(index, item) {
			temp = append(temp, item)
		}
	}
	return temp
}
