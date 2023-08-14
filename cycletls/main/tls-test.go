package main

import (
	"github.com/ChengHoward/CycleTLS/cycletls"
	"io"
	"log"
)

func main() {

	client := cycletls.Init()

	h2s := &cycletls.H2Settings{
		Settings: map[string]int{
			"HEADER_TABLE_SIZE":      65536,
			"ENABLE_PUSH":            0,
			"MAX_CONCURRENT_STREAMS": 1000,
			"INITIAL_WINDOW_SIZE":    6291456,
			"MAX_HEADER_LIST_SIZE":   262144,
			//"MAX_FRAME_SIZE":         16384,
		},
		SettingsOrder: []string{
			"HEADER_TABLE_SIZE",
			"ENABLE_PUSH",
			"MAX_CONCURRENT_STREAMS",
			"INITIAL_WINDOW_SIZE",
			"MAX_HEADER_LIST_SIZE",
		},
		ConnectionFlow: 15663105,
		HeaderPriority: map[string]interface{}{
			"weight":    256,
			"streamDep": 0,
			"exclusive": true,
		},
		PriorityFrames: []map[string]interface{}{
			{
				"streamID": 0,
				"priorityParam": map[string]interface{}{
					"weight":    0,
					"streamDep": 0,
					"exclusive": true,
				},
			},
			/*{
				"streamID": 3,
				"priorityParam": map[string]interface{}{
					"weight":    201,
					"streamDep": 0,
					"exclusive": false,
				},
			},*/
			/*{
				"streamID": 5,
				"priorityParam": map[string]interface{}{
					"weight":    101,
					"streamDep": 0,
					"exclusive": false,
				},
			},
			{
				"streamID": 7,
				"priorityParam": map[string]interface{}{
					"weight":    1,
					"streamDep": 0,
					"exclusive": false,
				},
			},
			{
				"streamID": 9,
				"priorityParam": map[string]interface{}{
					"weight":    1,
					"streamDep": 7,
					"exclusive": false,
				},
			},
			{
				"streamID": 11,
				"priorityParam": map[string]interface{}{
					"weight":    1,
					"streamDep": 3,
					"exclusive": false,
				},
			},
			{
				"streamID": 13,
				"priorityParam": map[string]interface{}{
					"weight":    241,
					"streamDep": 0,
					"exclusive": false,
				},
			},*/
		},
	}
	h2ss := cycletls.ToHTTP2Settings(h2s)

	resp, err := client.Do("https://tls.browserleaks.com/json", cycletls.Options{
		//resp, err := client.Do("https://tls.peet.ws/api/all", cycletls.Options{
		Ja3:           "771,4865-4866-4867-49195-49199-49196-49200-52393-52392-49171-49172-156-157-47-53,35-17513-16-51-23-43-11-0-5-27-65281-13-10-45-18-21,29-23-24,0",
		HTTP2Settings: h2ss,
		UserAgent:     "Mozilla/5.0 (X11; Ubuntu; Linux x86_64; rv:87.0) Gecko/20100101 Firefox/87.0",
		//Proxy:     "socks5://abc:321@127.0.0.1:1087",
		Headers: map[string]string{
			"Accept":          "Application/json, text/plain, */*",
			"Accept-Language": "en-US,en;q=0.5",
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
