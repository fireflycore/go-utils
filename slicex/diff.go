package slicex

// DiffSlice 比较两个切片的差异
// old: 旧切片数据
// new: 新切片数据
// 返回:
//   - add: 在 new 中存在但 old 中不存在的元素（新增）
//   - remove: 在 old 中存在但 new 中不存在的元素（删除）
//
// 注意：T 必须是 comparable 类型
func DiffSlice[T comparable](old, new []T) (add, remove []T) {
	oldSet := make(map[T]struct{}, len(old))
	newSet := make(map[T]struct{}, len(new))

	for _, item := range old {
		oldSet[item] = struct{}{}
	}

	for _, item := range new {
		newSet[item] = struct{}{}
	}

	addSet := make(map[T]struct{})
	for _, item := range new {
		if _, exists := oldSet[item]; exists {
			continue
		}
		if _, existed := addSet[item]; existed {
			continue
		}
		addSet[item] = struct{}{}
		add = append(add, item)
	}

	removeSet := make(map[T]struct{})
	for _, item := range old {
		if _, exists := newSet[item]; exists {
			continue
		}
		if _, existed := removeSet[item]; existed {
			continue
		}
		removeSet[item] = struct{}{}
		remove = append(remove, item)
	}

	return add, remove
}
