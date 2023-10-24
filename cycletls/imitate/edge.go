package imitate

import (
	"github.com/ChengHoward/CycleTLS/cycletls"
)

func Edge(options *cycletls.Options) {
	Chrome(options)

	options.Headers["Sec-Ch-Ua"] = `"Not/A)Brand";v="99", "Microsoft Edge";v="115", "Chromium";v="115"`
	options.UserAgent = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/115.0.0.0 Safari/537.36 Edg/115.0.1901.203"
	if options.Headers["Accept"] == "" {
		options.Headers["Accept"] = "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.7"
	}

}
