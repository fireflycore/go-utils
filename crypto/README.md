## Crypto (加密)

提供常用加解密算法的封装。

### 接口定义

```go
type Crypto interface {
    Encrypt(data []byte, key []byte) ([]byte, error)
    Decrypt(data []byte, key []byte) ([]byte, error)
}
```

### 支持的算法

#### AES (AES-GCM)

使用 AES-GCM 模式进行对称加密。

- **NewAESCrypto**
  - 签名: `func NewAESCrypto() Crypto`
  - 描述: 创建 AES 加密器。
  - 密钥要求: 16, 24, 或 32 字节。
  - 输出格式: Nonce + Ciphertext。

#### RSA (RSA-OAEP)

使用 RSA-OAEP (SHA256) 进行非对称加密。

- **NewRSACrypto**
  - 签名: `func NewRSACrypto() Crypto`
  - 描述: 创建 RSA 加密器。
  - 密钥要求: PEM 格式的公钥/私钥字节数据。

#### MD5

提供 MD5 哈希计算（注意：MD5 不实现 `Crypto` 接口）。

- **NewMD5**
  - 签名: `func NewMD5() *MD5`
- **Encrypt**
  - 签名: `func (ist *MD5) Encrypt(input, key string) string`
  - 描述: 计算 `key + input` 的 MD5 值，返回 hex 字符串。
- **Compare**
  - 签名: `func (ist *MD5) Compare(input, salt, hash string) bool`
  - 描述: 验证哈希匹配。

### 示例

```go
import "github.com/fireflycore/go-utils/crypto"

func main() {
    // AES 示例
    aes := crypto.NewAESCrypto()
    key := []byte("1234567890123456") // 16 bytes
    data := []byte("secret")
    encrypted, _ := aes.Encrypt(data, key)
    decrypted, _ := aes.Decrypt(encrypted, key)
}
```
