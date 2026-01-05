package file

import (
	"fmt"
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
	defer res.Body.Close()

	if res.StatusCode < http.StatusOK || res.StatusCode >= http.StatusMultipleChoices {
		return nil, fmt.Errorf("读取远程文件失败: status=%s", res.Status)
	}

	bytes, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	return bytes, nil
}
