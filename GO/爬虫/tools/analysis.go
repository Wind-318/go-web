package tools

import (
	"regexp"
)

// 得到正则表达式分析的页面结果
func RegexpHtml(url string, regexpRule string) []string {
	html, err := GetRequest(url)
	if err != nil {
		panic(err)
	}
	obj := regexp.MustCompile(regexpRule)
	arr := obj.FindAllStringSubmatch(string(html), -1)
	ret := make([]string, 0)
	for _, strs := range arr {
		ret = append(ret, strs[1])
	}
	return ret
}
