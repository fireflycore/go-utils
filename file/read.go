package file

import (
	"io"
	"net/http"
)

// ReadRemoteFile 读取远程文件内容
// url: 远程文件的 URL 地址
// 返回: 文件内容的字节切片。如果下载失败或读取失败，返回 error
// 注意：此函数会将整个文件加载到内存，不适合读取大文件
func ReadRemoteFile(url string) ([]byte, error) {
	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer func(Body io.ReadCloser) {
		if ce := Body.Close(); ce != nil {
			err = ce
		}
	}(res.Body)

	bytes, _ := io.ReadAll(res.Body)
	return bytes, err
}
