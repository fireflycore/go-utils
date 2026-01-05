package compress

// Compress 定义数据压缩与解压的通用接口
// 实现该接口的类型应当提供线程安全的压缩和解压方法
type Compress interface {
	// Compress 将原始数据压缩为字节切片
	// 如果压缩过程中发生错误，返回 error
	Compress(data []byte) ([]byte, error)

	// Decompress 将压缩后的数据解压为原始字节切片
	// 如果解压过程中发生错误（如数据损坏），返回 error
	Decompress(data []byte) ([]byte, error)
}
