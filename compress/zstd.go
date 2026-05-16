package compress

import (
	"bytes"
	"io"

	"github.com/klauspost/compress/zstd"
)

// Zstd 实现 Compress 接口，提供基于 zstd 的压缩与解压能力。
// 注意：该实现会将输入与输出完整加载到内存中，不适合超大数据。
type Zstd struct {
	encoder *zstd.Encoder
	decoder *zstd.Decoder
}

// NewZstd 创建一个新的 Zstd 压缩器实例，使用默认压缩级别。
// 若初始化失败（极少情况），会 panic，可依据需要改为返回 error。
func NewZstd() Compress {
	enc, err := zstd.NewWriter(nil, zstd.WithEncoderLevel(zstd.SpeedDefault))
	if err != nil {
		panic(err)
	}
	dec, err := zstd.NewReader(nil)
	if err != nil {
		panic(err)
	}
	return &Zstd{
		encoder: enc,
		decoder: dec,
	}
}

// Compress 使用 zstd 压缩数据
func (z *Zstd) Compress(data []byte) ([]byte, error) {
	var buf bytes.Buffer
	// 重置 encoder 并绑定到新的输出
	z.encoder.Reset(&buf)
	if _, err := z.encoder.Write(data); err != nil {
		return nil, err
	}
	if err := z.encoder.Close(); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

// Decompress 使用 zstd 解压数据
// data: zstd 格式的压缩数据
// 返回解压后的原始数据，如果数据格式不正确或读取失败则返回错误
func (z *Zstd) Decompress(data []byte) ([]byte, error) {
	// 重置 decoder 并绑定到新的输入
	if err := z.decoder.Reset(bytes.NewReader(data)); err != nil {
		return nil, err
	}
	defer z.decoder.Reset(nil) // 解压完成后释放内部引用
	return io.ReadAll(z.decoder)
}
