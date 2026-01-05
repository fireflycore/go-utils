package compress

import (
	"bytes"
	"compress/gzip"
	"io"
)

// GZIP 实现 Compress 接口，提供基于 gzip 的压缩与解压能力。
// 注意：该实现会将输入与输出完整加载到内存中，不适合超大数据。
type GZIP struct{}

// NewGZIP 创建一个新的 GZIP 压缩器实例。
func NewGZIP() Compress {
	return &GZIP{}
}

// Compress 压缩数据
func (g *GZIP) Compress(data []byte) ([]byte, error) {
	var buf bytes.Buffer
	gz := gzip.NewWriter(&buf)
	if _, err := gz.Write(data); err != nil {
		return nil, err
	}
	if err := gz.Close(); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

// Decompress 使用 gzip 算法解压数据
// data: gzip 格式的压缩数据
// 返回解压后的原始数据，如果数据格式不正确或读取失败则返回错误
func (g *GZIP) Decompress(data []byte) ([]byte, error) {
	gz, err := gzip.NewReader(bytes.NewReader(data))
	if err != nil {
		return nil, err
	}
	defer gz.Close()
	return io.ReadAll(gz)
}
