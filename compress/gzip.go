package compress

import (
	"bytes"
	"compress/gzip"
	"io"
)

type GZIP struct {
}

func NewGZIP() Compress {
	return &GZIP{}
}

// Compress 压缩数据
func (ist *GZIP) Compress(data []byte) ([]byte, error) {
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
func (ist *GZIP) Decompress(data []byte) ([]byte, error) {
	gz, err := gzip.NewReader(bytes.NewReader(data))
	if err != nil {
		return nil, err
	}
	defer gz.Close()
	return io.ReadAll(gz)
}
