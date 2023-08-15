package imitate

import (
	"github.com/ChengHoward/CycleTLS/cycletls"
)

var firefoxH2Settings = &cycletls.H2Settings{
	Settings: map[string]int{
		"HEADER_TABLE_SIZE":      65536,
		"ENABLE_PUSH":            0,
		"MAX_CONCURRENT_STREAMS": 1000,
		"INITIAL_WINDOW_SIZE":    131072,
		"MAX_HEADER_LIST_SIZE":   262144,
		"MAX_FRAME_SIZE":         16384,
	},
	SettingsOrder: []string{
		"HEADER_TABLE_SIZE",
		"INITIAL_WINDOW_SIZE",
		"MAX_FRAME_SIZE",
	},
	ConnectionFlow: 12517377,
	HeaderPriority: map[string]interface{}{
		"weight":    256,
		"streamDep": 0,
		"exclusive": true,
	},
	PriorityFrames: []map[string]interface{}{
		/*{
			"streamID": 0,
			"priorityParam": map[string]interface{}{
				"weight":    0,
				"streamDep": 0,
				"exclusive": true,
			},
		},*/
		{
			"streamID": 3,
			"priorityParam": map[string]interface{}{
				"weight":    201,
				"streamDep": 0,
				"exclusive": false,
			},
		},
		{
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
		},
	},
}
var firefoxHttp2Setting = cycletls.ToHTTP2Settings(firefoxH2Settings)

func Firefox(options *cycletls.Options) {
	//"0-23-65281-10-11-35-16-5-51-43-13-45-28-21"
	//"0-23-65281-10-11-35-16-5-34-51-43-13-45-28-21"
	options.Ja3 = "771,4865-4867-4866-49195-49199-52393-52392-49196-49200-49162-49161-49171-49172-156-157-47-53,0-23-65281-10-11-35-16-5-34-51-43-13-45-28-21,29-23-24-25-256-257,0"
	options.HTTP2Settings = firefoxHttp2Setting
	options.PHeaderOrderKeys = []string{
		":method",
		":path",
		":authority",
		":scheme",
	}
}
