package array

// DiffSlice 比较两个切片的差异
// old: 旧切片数据
// new: 新切片数据
// 返回:
//   - add: 在 new 中存在但 old 中不存在的元素（新增）
//   - remove: 在 old 中存在但 new 中不存在的元素（删除）
//
// 注意：T 必须是 comparable 类型
func DiffSlice[T comparable](old, new []T) (add, remove []T) {
	// 使用双map实现O(1)存在性检查 + 自动去重
	oldSet := make(map[T]struct{}, len(old))
	newSet := make(map[T]struct{}, len(new))

	// 初始化旧数据集（保持原始顺序用于删除检测）
	for _, item := range old {
		oldSet[item] = struct{}{}
	}

	// 初始化新数据集（保持原始顺序用于新增检测）
	for _, item := range new {
		newSet[item] = struct{}{}
	}

	// 检测新增项（使用新集合顺序保证确定性）
	add = make([]T, 0, len(newSet))
	for item := range newSet {
		if _, exists := oldSet[item]; !exists {
			add = append(add, item)
		}
	}

	// 检测删除项（使用旧集合顺序保证确定性）
	remove = make([]T, 0, len(oldSet))
	for item := range oldSet {
		if _, exists := newSet[item]; !exists {
			remove = append(remove, item)
		}
	}

	return add, remove
}
