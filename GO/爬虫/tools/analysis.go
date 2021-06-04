package tools

import (
	"regexp"
)

func AnalysisHtml(html []byte, regexpRule string) []string {
	obj := regexp.MustCompile(regexpRule)
	arr := obj.FindAllStringSubmatch(string(html), -1)
	ret := make([]string, 0)
	for _, strs := range arr {
		ret = append(ret, strs[1])
	}
	return ret
}
