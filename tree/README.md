## Tree（树）

## 1、提供标签树（Label Tree）路径操作工具。
- `GetLTreeDepth(path string) int`: 计算路径深度。
- `BuildLTreePath(parentPath, currentValue string) string`: 构建路径。
- `ExtractLastSegment(path string) string`: 提取路径最后一段。
- `IsDescendantPath(descendantPath, ancestorPath string) bool`: 判断是否为后代路径。
- `ReplacePathPrefix(oldPath, newPath, fullPath string) string`: 替换路径前缀。
- `ReplaceCurrentPath(oldPath, currentPath string) string`: 替换当前路径节点。