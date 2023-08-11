package main

import (
	"fmt"
	"github.com/ChengHoward/CycleTLS/cycletls"
	"io"
)

func main() {
	fileURL := "https://194.9.149.45/zh/online-business/schedule/interactive-schedule/interactive-schedule-solution.html?sn=TIANJIN%20XINGANG&sl=CNTXG&sp=&en=ENSENADA,%20BCN&el=MXESE&ep=22800&exportHaulage=MH&importHaulage=MH&departureDate=2023-08-04&weeksAfterStart=4&reefer=N&exportMot=VE&importMot=VE"
	client := cycletls.Init()

	resp, err := client.Do(fileURL, cycletls.Options{
		Body:      nil,
		Ja3:       "771,4865-4866-4867-49195-49199-49196-49200-52393-52392-49171-49172-156-157-47-53,0-5-10-11-13-16-18-21-23-27-35-43-45-51-17513-65281,29-23-24,0",
		UserAgent: "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/114.0.0.0 Safari/537.36",
		//Stream:          true,
		Timeout:         120,
		DisableRedirect: true,
		Proxy:           "socks5://127.0.0.1:10808",
		Headers: map[string]string{
			"Host": "www.hapag-lloyd.com",
			//"Connection":                "Upgrade, HTTP2-Settings",
			//"Upgrade":                   "h2c",
			//"Http2-Settings":            "AAEAAQAAAAIAAAAAAAMAAAPoAAQAYAAAAAYABAAA",
			"Sec-Ch-Ua":                 "\"Chromium\";v=\"110\", \"Not A(Brand\";v=\"24\", \"Google Chrome\";v=\"110\"",
			"Sec-Ch-Ua-Mobile":          "?0",
			"Sec-Ch-Ua-Platform":        "\"Windows\"",
			"Upgrade-Insecure-Requests": "1",
			"User-Agent":                "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/110.0.0.0 Safari/537.36",
			"Accept":                    "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.7",
			"Sec-Fetch-Site":            "none",
			"Sec-Fetch-Mode":            "navigate",
			"Sec-Fetch-User":            "?1",
			"Sec-Fetch-Dest":            "document",
			"Accept-Encoding":           "gzip, deflate, br",
			"Accept-Language":           "en-US,en;q=0.9",
			"Pragma":                    "no-cache",
		},
	}, "GET")
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(resp.Body)

	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println(resp.Status)
	fmt.Println(string(body))
}
