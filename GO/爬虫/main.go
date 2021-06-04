package main

import (
	"GoProject/GoWeb/GO/爬虫/tools"
	"fmt"
)

func main() {
	root := "https://finance.sina.com.cn/stock/"
	html, _ := tools.GetRequest(root)
	regexpRule := `<a target="_blank" href="(https://finance.sina.com.cn/stock/[\S]+?)">`
	arr := tools.AnalysisHtml(html, regexpRule)

	for _, url := range arr {
		html, _ = tools.GetRequest(url)
		//titleRule := `<title>([\s\S]+?)</title>`
		//title := tools.AnalysisHtml(html, titleRule)[0]
		//os.Mkdir(`.\news\`+title, 0644)

		sourceRule := `"(//[\s\S]+?.js[\S]*?)"`
		sources := tools.AnalysisHtml(html, sourceRule)
		for _, i := range sources {
			fmt.Println(i)
			break
		}
		break
	}
}
