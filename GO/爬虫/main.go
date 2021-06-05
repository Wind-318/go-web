package main

import (
	"GoProject/GoWeb/GO/爬虫/tools"
)

func main() {
	root := "https://finance.sina.com.cn/stock/"
	regexpRule := `<a target="_blank" href="(https://finance.sina.com.cn/stock/[\S]+?)">`
	arr := tools.RegexpHtml(root, regexpRule)

	for _, url := range arr {
		titleRule := `<title>([\s\S]+?)</title>`
		title := tools.RegexpHtml(url, titleRule)[0]
		pos := ".\\news\\" + title + "\\"
		tools.DownloadHtmlSource(url, pos)
	}
}
