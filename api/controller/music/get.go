package music

import (
	"GoAPI/helper"
)

func musicGET(url string, params map[string]any, server string) (result helper.CurlResult) {

	var headers = make(map[string]any)

	switch server {
	case "netease":
		{
			headers["Referer"] = "https://music.163.com/"
			headers["cookie"] = "appver=8.2.30; os=iPhone OS; osver=15.0; EVNSM=1.0.0; buildver=2206; channel=distribution; machineid=iPhone13.3"
			headers["User-Agent"] = "Mozilla/5.0 (iPhone; CPU iPhone OS 15_0 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Mobile/15E148 CloudMusic/0.1.1 NeteaseMusic/8.2.30"
		}
	default:
		{
		}
	}

	return helper.HTTP.GET(url, params, headers)
}
