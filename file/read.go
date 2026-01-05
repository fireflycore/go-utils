package file

import (
	"io"
	"net/http"
)

// ReadRemoteFile 读取远程文件内容
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
