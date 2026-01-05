package compress

// Compress 定义压缩接口
type Compress interface {
	Compress(data []byte) ([]byte, error)
	Decompress(data []byte) ([]byte, error)
}
