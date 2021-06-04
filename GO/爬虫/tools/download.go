package tools

import (
	"io/ioutil"
	"math/rand"
	"net/http"
	"strconv"
)

func DownloadGet(url string, position ...string) error {
	if len(position) > 1 {
		panic("too many args")
	}
	pos := ".\\" + strconv.Itoa(rand.Intn(100000))
	if len(position) == 1 {
		pos = position[0]
	}

	res, err := http.Get(url)
	if err != nil {
		return err
	}
	defer res.Body.Close()
	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}

	ioutil.WriteFile(pos, data, 0644)
	return nil
}
