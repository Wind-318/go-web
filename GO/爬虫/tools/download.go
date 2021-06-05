package tools

import (
	"io/ioutil"
	"math/rand"
	"net/http"
	"os"
	"regexp"
	"strconv"
)

// 下载资源
func DownloadGet(url string, position ...string) error {
	if len(position) > 1 {
		panic("too many args")
	}
	pos := ".\\files\\" + strconv.Itoa(rand.Intn(100000))
	if len(position) == 1 {
		pos = position[0]
	}

	// 去除不合法字符
	obj := regexp.MustCompile(`[/\:*?"<>|]`)
	pos = obj.ReplaceAllString(pos, "")

	// 正则表达式搜索文件所需文件夹是否存在
	obj = regexp.MustCompile(`([\s\S]+)\\`)
	fileroot := obj.FindAllStringSubmatch(pos, -1)[0][1]

	// 递归创建所有所需文件夹
	os.MkdirAll(fileroot, 0644)

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

// 下载 html 页面中所有资源
func DownloadHtmlSource(url string, position ...string) {
	if len(position) > 1 {
		panic("too many args")
	}
	pos := ".\\files\\" + strconv.Itoa(rand.Intn(100000))
	if len(position) == 1 {
		pos = position[0]
	}

	// 资源文件匹配规则
	sourceRule := `(//[\S]+?((\.css)|(\.jpg)|(\.png)|(\.js))[\S]*?)`
	sources := RegexpHtml(url, sourceRule)
	for _, urls := range sources {
		DownloadGet("https:"+urls, pos+urls)
	}
	DownloadGet(url, pos+"1.html")
}
