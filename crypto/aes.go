package crypto

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"errors"
	"io"
)

// AES 实现了 Crypto 接口，提供基于 AES-GCM 模式的加解密功能
// 使用 Galois/Counter Mode (GCM) 提供认证加密
type AES struct{}

// NewAESCrypto 创建一个新的 AES 加密器实例
func NewAESCrypto() Crypto {
	return &AES{}
}

// Encrypt 使用 AES-GCM 算法加密数据
// data: 待加密的明文
// key: AES 密钥，长度必须为 16, 24 或 32 字节（对应 AES-128, AES-192, AES-256）
// 返回格式: [nonce(随机值) + ciphertext(密文)]
func (e *AES) Encrypt(data []byte, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	nonce := make([]byte, gcm.NonceSize())
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		return nil, err
	}

	return gcm.Seal(nonce, nonce, data, nil), nil
}

// Decrypt 使用 AES-GCM 算法解密数据
// data: 密文数据，格式必须为 [nonce + ciphertext]
// key: 解密密钥，必须与加密时使用的密钥一致
func (e *AES) Decrypt(data []byte, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	nonceSize := gcm.NonceSize()
	if len(data) < nonceSize {
		return nil, errors.New("ciphertext too short")
	}

	nonce, ciphertext := data[:nonceSize], data[nonceSize:]
	return gcm.Open(nil, nonce, ciphertext, nil)
}
