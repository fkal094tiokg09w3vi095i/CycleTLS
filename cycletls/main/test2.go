package main

import (
	"fmt"
	"github.com/ChengHoward/CycleTLS/cycletls"
	"github.com/ChengHoward/CycleTLS/cycletls/imitate"
	"io"
)

func main() {
	URL := "http://www.26ks.org/book/4471/62835032.html"
	//URL := "http://ipinfo.io/ip"
	client := cycletls.Init()

	options := cycletls.Options{
		Timeout: 120,
		Headers: map[string]string{
			"Accept": "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.7",
		},
		Proxy: "http://127.0.0.1:10809",
		//Proxy: "http://127.0.0.1:60010",
		//Proxy: "http://127.0.0.1:8999",
		//Proxy: "http://howard:pppppppp@47.242.250.113:10087",
	}
	imitate.Chrome(&options)
	for s, s2 := range options.Headers {
		fmt.Println(s, ": ", s2)
	}
	resp, err := client.Do(URL, options, "GET")
	defer resp.Body.Close()

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
