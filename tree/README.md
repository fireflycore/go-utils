## Tree (树)

提供 PostgreSQL ltree 风格的标签树路径操作工具。

### API

#### LTree

用于处理类似 "A.B.C" 的层级路径结构。

- **NewLTree**
  - 签名: `func NewLTree() *LTree`
  - 描述: 创建 LTree 工具实例。

#### Methods

- **GetLTreeDepth**
  - 签名: `func (t *LTree) GetLTreeDepth(path string) int`
  - 描述: 计算路径深度（即节点数量）。
- **BuildLTreePath**
  - 签名: `func (t *LTree) BuildLTreePath(parentPath, currentValue string) string`
  - 描述: 构建子节点路径。
- **ExtractLastSegment**
  - 签名: `func (t *LTree) ExtractLastSegment(path string) string`
  - 描述: 提取路径最后一段（当前节点名）。
- **IsDescendantPath**
  - 签名: `func (t *LTree) IsDescendantPath(descendantPath, ancestorPath string) bool`
  - 描述: 判断是否为后代路径。
- **ReplacePathPrefix**
  - 签名: `func (t *LTree) ReplacePathPrefix(oldPath, newPath, fullPath string) string`
  - 描述: 替换路径前缀（用于移动子树）。
- **ReplaceCurrentPath**
  - 签名: `func (t *LTree) ReplaceCurrentPath(oldPath, currentPath string) string`
  - 描述: 重命名当前节点。

### 示例

```go
import (
    "fmt"
    "github.com/fireflycore/go-utils/tree"
)

func main() {
    lt := tree.NewLTree()
    
    path := "Region.China.Beijing"
    fmt.Println(lt.GetLTreeDepth(path)) // 3
    fmt.Println(lt.ExtractLastSegment(path)) // Beijing
    
    newPath := lt.ReplacePathPrefix("Region.China", "Area.Asia", path)
    fmt.Println(newPath) // Area.Asia.Beijing
}
```
