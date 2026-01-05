## File (文件)
提供文件读写辅助功能。

- **ReadRemoteFile**
  - 签名: `func ReadRemoteFile(url string) ([]byte, error)`
  - 描述: 读取远程 URL 的文件内容。
- **WriteLocalFile**
  - 签名: `func WriteLocalFile(path string, bytes []byte) error`
  - 描述: 将字节内容写入本地文件，如果目录不存在会自动创建。
