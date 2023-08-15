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
	// 771,4865-4867-4866-49195-49199-52393-52392-49196-49200-49162-49161-49171-49172-51-57-47-53-10,0-23-65281-10-11-35-16-5-51-43-13-45-28-21,29-23-24-25-256-257,0
	// 771,4865-4866-4867-49195-49199-49196-49200-52393-52392-49171-49172-156-157-47-53,35-17513-16-51-23-43-11-0-5-27-65281-13-10-45-18-21,29-23-24,0
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
}
