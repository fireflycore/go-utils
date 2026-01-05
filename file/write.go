package file

import (
	"os"
	"path/filepath"
)

// WriteLocalFile 将内容写入本地文件
// 如果目录不存在，会自动创建
func WriteLocalFile(path string, bytes []byte) error {
	// 分割路径，得到目录部分和文件名部分
	dir, _ := filepath.Split(path)

	if err := os.MkdirAll(dir, 0755); err != nil {
		return err
	}

	if err := os.WriteFile(path, bytes, 0666); err != nil {
		return err
	}

	return nil
}
