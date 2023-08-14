package main

import (
	"github.com/ChengHoward/CycleTLS/cycletls"
	"io"
	"log"
)

func main() {

	client := cycletls.Init()
	resp, err := client.Do("https://tls.browserleaks.com/json", cycletls.Options{
		Ja3:       "771,4865-4866-4867-49195-49199-49196-49200-52393-52392-49171-49172-156-157-47-53,35-17513-16-51-23-43-11-0-5-27-65281-13-10-45-18-21,29-23-24,0",
		UserAgent: "Mozilla/5.0 (X11; Ubuntu; Linux x86_64; rv:87.0) Gecko/20100101 Firefox/87.0",
		//Proxy:     "socks5://abc:321@127.0.0.1:1087",
		Headers: map[string]string{
			"Accept": "Application/json, text/plain, */*",
		},
	}, "GET")
	if err != nil {
		log.Print("Request Failed: " + err.Error())
		return
	}
	body, err := io.ReadAll(resp.Body)

	log.Print(resp.Status)
	log.Print(string(body))

}
