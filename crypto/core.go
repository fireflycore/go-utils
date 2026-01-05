package crypto

// Crypto 定义通用加解密接口
// 注意：不同的实现可能对 data 和 key 有不同的长度或格式要求
type Crypto interface {
	// Encrypt 加密数据
	// data: 原始数据
	// key: 加密密钥
	// 返回加密后的数据，若出错返回 error
	Encrypt(data []byte, key []byte) ([]byte, error)

	// Decrypt 解密数据
	// data: 密文数据
	// key: 解密密钥
	// 返回解密后的原始数据，若出错返回 error
	Decrypt(data []byte, key []byte) ([]byte, error)
}
