package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"regexp"
	"time"
)

func main() {
	res, _ := http.Get("https://www.zhihu.com/question/339135205")
	defer res.Body.Close()

	html, _ := ioutil.ReadAll(res.Body)

	// 获取一个正则表达对象
	obj := regexp.MustCompile(`<a href="([\s\S]+?)"[\s\S]+?>([\s\S]+?)</a>`)
	// 得到字符串
	arr := obj.FindAllStringSubmatch(string(html), -1)

	target := make([]string, 0)
	name := make([]string, 0)

	for i, n := 3, len(arr); i < n; i++ {
		target = append(target, arr[i][1][66:72])
		name = append(name, arr[i][2])
	}

	err := os.Mkdir(`.\面经`, os.ModePerm)
	if err != nil {
		fmt.Println(err)
	}

	for index, val := range target {
		url := "https://www.nowcoder.com/discuss/" + val + "?from=zhnkw"
		url += "1"
		objs := regexp.MustCompile(`[\/:*?"<>|]`)
		name[index] = objs.ReplaceAllString(name[index], "")

		get(url, name[index])
	}
}

func get(url, name string) {
	clients := &http.Client{
		Timeout: 3 * time.Second,
	}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println(err)
		return
	}

	res, err := clients.Do(req)
	defer res.Body.Close()
	if err != nil {
		fmt.Println(err)
		return
	}

	html, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}

	err = ioutil.WriteFile(`.\面经\`+name+`.html`, html, 0644)
	if err != nil {
		fmt.Println(err)
	}
}
