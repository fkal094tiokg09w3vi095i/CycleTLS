package imitate

import (
	"github.com/ChengHoward/CycleTLS/cycletls"
	"math/rand"
	"strings"
)

const chromeExtension = "0-5-10-11-13-16-18-21-23-27-35-43-45-51-17513-65281"

var chromeH2Settings = &cycletls.H2Settings{
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
var chromeHttp2Setting = cycletls.ToHTTP2Settings(chromeH2Settings)

func shuffleExtension() string {
	extension := strings.Split(chromeExtension, "-")
	for i := range extension {
		j := rand.Intn(len(extension))
		extension[i], extension[j] = extension[j], extension[i]
	}
	return strings.Join(extension, "-")
}

func chrome109Ja3() string {
	return "771,4865-4866-4867-49195-49199-49196-49200-52393-52392-49171-49172-156-157-47-53" + "," + shuffleExtension() + ",29-23-24,0"
}

func Chrome(options *cycletls.Options) {
	options.Ja3 = chrome109Ja3()
	options.HTTP2Settings = chromeHttp2Setting
	options.PHeaderOrderKeys = []string{
		":method",
		":authority",
		":scheme",
		":path",
	}
	if options.Headers == nil {
		options.Headers = make(map[string]string)
	}

	options.Headers["Sec-Ch-Ua"] = `"Not.A/Brand";v="8", "Chromium";v="114", "Google Chrome";v="114"`
	options.Headers["Sec-Ch-Ua-Mobile"] = "?0"
	options.Headers["Sec-Ch-Ua-Platform"] = `"Windows"`
	options.Headers["Sec-Fetch-Dest"] = "document"
	options.Headers["Sec-Fetch-Mode"] = "navigate"
	options.Headers["Sec-Fetch-Site"] = "none"
	options.Headers["Sec-Fetch-User"] = "?1"
	options.UserAgent = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/114.0.0.0 Safari/537.36"

}
