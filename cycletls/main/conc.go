package main

import (
	"fmt"
	"github.com/ChengHoward/CycleTLS/cycletls"
	"io"
	"sync"
	"time"
)

func main() {
	URL := "http://api.example.io/"
	client := cycletls.Init()
	wg := new(sync.WaitGroup)
	// 开始时间
	t1 := time.Now()

	// 循环一万次
	for i := 0; i < 1000; i++ {
		go func(item int) {
			resp, err := client.Do(URL, cycletls.Options{
				Body:      nil,
				Ja3:       "771,4865-4866-4867-49195-49199-49196-49200-52393-52392-49171-49172-156-157-47-53,0-5-10-11-13-16-18-21-23-27-35-43-45-51-17513-65281,29-23-24,0",
				UserAgent: "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/114.0.0.0 Safari/537.36",
				Stream:    true,
				Timeout:   120,
				Headers: map[string]string{
					"x-cb-apikey":   "30e9707df33941e3a5f64f2ea4a19186",
					"x-cb-host":     "192.168.31.111",
					"x-cb-protocol": "http",
					"x-cb-options":  "direct",
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
			fmt.Println(string(body), item)
			wg.Done()
		}(i)
		wg.Add(1)
	}
	wg.Wait()
	// 结束时间
	elapsed := time.Since(t1)
	fmt.Println("App elapsed: ", elapsed)
}
