package main

import (
	"fmt"
	"github.com/ChengHoward/CycleTLS/cycletls"
	"github.com/ChengHoward/CycleTLS/cycletls/imitate"
	"io"
)

func main() {
	URL := "https://k.apiairasia.com/checkin/auth/getssrdata"
	client := cycletls.Init()

	options := cycletls.Options{
		Body:    "{}",
		Timeout: 120,
		Headers: map[string]string{
			"Accept":                    "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.7",
			"upgrade-insecure-requests": "1",
			"Referer":                   "https://www.airasia.com/",
			"Origin":                    "https://www.airasia.com",
			"Content-Type":              "text/plain",
			"Channel_hash":              "ee47af6796903d24cce0c8e12d78963e3c06c16316ef55a88568065e",
		},
		//Proxy: "http://192.168.31.192:10087",
		Proxy: "http://127.0.0.1:60010",
	}
	imitate.Chrome(&options)
	resp, err := client.Do(URL, &options, "POST")
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
