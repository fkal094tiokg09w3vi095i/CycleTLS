package main

import (
	"github.com/ChengHoward/CycleTLS/cycletls"
	"github.com/ChengHoward/CycleTLS/cycletls/imitate"
	"io"
	"log"
)

func main() {

	client := cycletls.Init()

	options := cycletls.Options{
		Timeout: 60,
		Headers: map[string]string{
			"sec-fetch-site":  "same-site",
			"sec-fetch-mode":  "cors",
			"sec-fetch-user":  "?1",
			"sec-fetch-dest":  "empty",
			"Accept":          "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.7",
			"Accept-Language": "zh-CN,zh;q=0.9,en;q=0.8,sq;q=0.7",
			"Referer":         "https://link.2505.top/article/10.1007/s10029-020-02328-x",
			"User-Agent":      "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/114.0.5735.134 Electron/25.2.0 Safari/537.36",
			"Host":            "www.vijesti.me",
		},
		Proxy:   "socks5://127.0.0.1:10808",
		Cookies: []cycletls.Cookie{},
	}

	imitate.Chrome(&options)
	//imitate.Firefox(&options)
	//imitate.Chromium(&options)
	//imitate.Edge(&options)
	//imitate.Chrome(&options)

	//options.UserAgent = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/114.0.5735.134 Electron/25.2.0 Safari/537.36"

	resp, err := client.Do("https://92.249.52.175", options, "GET")
	if err != nil {
		log.Print("Request Failed: " + err.Error())
		return
	}
	body, err := io.ReadAll(resp.Body)

	log.Print(resp.Status)
	log.Print(string(body))

}
