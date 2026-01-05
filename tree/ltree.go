package tree

import "strings"

// LTree 提供 PostgreSQL ltree 风格的标签树路径操作工具
// 用于处理类似 "A.B.C" 的层级路径结构
type LTree struct{}

// NewLTree 创建一个新的 LTree 工具实例
func NewLTree() *LTree {
	return &LTree{}
}

// GetLTreeDepth 计算标签树路径的深度
// 深度定义为路径中层级的数量（即点的数量 + 1）
// 例如: "A.B.C" 的深度为 3
func (t *LTree) GetLTreeDepth(path string) int {
	if path == "" {
		return 0
	}
	return strings.Count(path, ".") + 1
}

// BuildLTreePath 构建子节点路径
// parentPath: 父节点路径
// currentValue: 当前节点值
// 返回: "parentPath.currentValue"，如果 parentPath 为空则返回 currentValue
func (t *LTree) BuildLTreePath(parentPath, currentValue string) string {
	if parentPath == "" {
		return currentValue
	}
	return parentPath + "." + currentValue
}

// ExtractLastSegment 提取路径的最后一段（即当前节点名）
// 例如: "A.B.C" -> "C"
func (t *LTree) ExtractLastSegment(path string) string {
	parts := strings.Split(path, ".")
	if len(parts) == 0 {
		return ""
	}
	return parts[len(parts)-1]
}

// IsDescendantPath 检查路径是否是后代路径
func (t *LTree) IsDescendantPath(descendantPath, ancestorPath string) bool {
	if ancestorPath == "" {
		return true
	}
	return strings.HasPrefix(descendantPath, ancestorPath+".")
}

// ReplacePathPrefix 替换路径前缀，常用于移动子树
// oldPath: 旧的前缀（如 "A.B"）
// newPath: 新的前缀（如 "X.Y"）
// fullPath: 包含旧前缀的完整路径（如 "A.B.C.D"）
// 返回: 替换后的路径（如 "X.Y.C.D"）。如果 fullPath 不以 oldPath 开头，则原样返回
func (t *LTree) ReplacePathPrefix(oldPath, newPath, fullPath string) string {
	if !strings.HasPrefix(fullPath, oldPath+".") {
		return fullPath
	}
	return newPath + "." + strings.TrimPrefix(fullPath, oldPath+".")
}

// ReplaceCurrentPath 替换当前路径
func (t *LTree) ReplaceCurrentPath(oldPath, currentPath string) string {
	if oldPath == "" {
		return currentPath
	}
	parts := strings.Split(oldPath, ".")
	parts = parts[:len(parts)-1]
	parts = append(parts, currentPath)
	return strings.Join(parts, ".")
}
