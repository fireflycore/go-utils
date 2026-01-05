package crypto

import (
	"crypto/md5"
	"encoding/hex"
	"io"
)

// MD5 提供 MD5 哈希计算功能
// 注意：MD5 已被认为不再安全，仅建议用于非安全敏感场景（如文件完整性校验、简单的指纹生成）
// 该结构体未实现 Crypto 接口，因为其 Encrypt 方法签名不同
type MD5 struct{}

// NewMD5 创建一个新的 MD5 工具实例
func NewMD5() *MD5 {
	return &MD5{}
}

// Encrypt 计算字符串的 MD5 哈希值
// input: 待哈希的字符串
// key: 可选的盐值（salt），如果非空，将在 input 之前写入哈希计算器
// 返回: 32位十六进制字符串
func (m *MD5) Encrypt(input, key string) string {
	hash := md5.New()
	if key != "" {
		_, _ = io.WriteString(hash, key)
	}
	_, _ = io.WriteString(hash, input)
	return hex.EncodeToString(hash.Sum(nil))
}

// Compare 比较输入字符串的哈希值是否与给定哈希匹配
// input: 待验证的原始字符串
// salt: 计算哈希时使用的盐值
// hash: 预期的 MD5 哈希值（十六进制字符串）
func (m *MD5) Compare(input, salt, hash string) bool {
	return m.Encrypt(input, salt) == hash
}
