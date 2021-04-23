package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"time"
)

func main() {
	res, _ := http.Get("https://www.zhihu.com/question/339135205")
	defer res.Body.Close()

	html, _ := ioutil.ReadAll(res.Body)
	obj := regexp.MustCompile(`<a href="([\s\S]+?)"`)
	arr := obj.FindAllStringSubmatch(string(html), -1)

	target := make([]string, 0)

	for i, n := 3, len(arr); i < n; i++ {
		target = append(target, arr[i][1][66:72])
	}

	for _, val := range target {
		url := "https://www.nowcoder.com/discuss/" + val + "?from=zhnkw"
		get(url)
		break
	}

}

func get(url string) {
	clients := &http.Client{
		Timeout: 3 * time.Second,
	}
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/90.0.4430.72 Safari/537.36 Edg/90.0.818.42")
	req.Header.Set("Cookie", "NOWCODERUID=6C11A31BEB8625EA9A83EBE6EE356B9F; NOWCODERCLINETID=9F8EFCFAE485B69288D6A54219BC06A4; NOWCODERQUERY=%E5%AD%97%E8%8A%82%E8%B7%B3%E5%8A%A8; __snaker__id=NnFSyErNvm7Ps6hU; _9755xjdesxxd_=32; gdxidpyhxdE=EVZedvrkOkvUgOBebjUWmz760cl7AEpebvkApdcP4OG6%5CxCjxeOIBQaL%2BQu8ZtxZcW5Vpvo2ZUx%2BGttaqgRUumsBtqSajy8MSz%2FDvazd7v7%5C27%5Cnlh%5CS7r3bNYzMC4zraS122wLGAWf%2FYkAmbAYT%2B%5CzBipimK%2F6kBy7veKw2D6A%2BLsRv%3A1617794728530; t=FD52644B4C8C1499016642020C5AE4CE; live_chapter_2968=Ad5OOlo2wbvpHttCktje; Hm_lpvt_a808a1326b6c06c437de769d1b85b870=1618186760; Hm_lvt_a808a1326b6c06c437de769d1b85b870=1618186760; gr_user_id=936bc1e6-f75e-42ef-bd1a-dc2b5d2fdf89; c196c3667d214851b11233f5c17f99d5_gr_last_sent_cs1=926255587; c196c3667d214851b11233f5c17f99d5_gr_cs1=926255587; from=zhnkw; SERVERID=f24cbffaf8c883b27da19f52fb8cda88|1619189661|1619189442")
	req.Header.Set("Referer", "https://link.zhihu.com/?target=https%3A//www.nowcoder.com/discuss/206520%3Ffrom%3Dzhnkw")

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

	ioutil.WriteFile("1.html", html, 0644)
}
