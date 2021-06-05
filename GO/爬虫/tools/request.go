package tools

import (
	"io/ioutil"
	"net/http"
)

// 获取请求 url 的 html 页面
func GetRequest(url string) ([]byte, error) {
	res, err := http.Get(url)
	defer res.Body.Close()
	if err != nil {
		return nil, err
	}
	html, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	return html, nil
}
