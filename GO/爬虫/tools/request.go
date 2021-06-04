package tools

import (
	"io/ioutil"
	"net/http"
)

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
