## Version (版本)
提供语义化版本管理工具。

- **Version** 结构体: 包含 `Major`, `Minor`, `Patch`。
- **NewVersion**
  - 签名: `func NewVersion(verStr string) (*Version, error)`
  - 描述: 解析版本字符串（如 "v1.0.1"）。
- **相关方法**:
  - `String()`: 格式化为字符串。
  - `Increment()`: 版本号自增逻辑（Patch -> Minor -> Major）。
  - `Compare(other *Version) int`: 版本比较。
