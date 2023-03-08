package service

import (
	"GinAPI/helper"
)


func GetList(id string, server string) string {
	switch server {
	case "netease":
		{
			params := make(map[string]any)
			params["id"] = id
			params["n"] = 1000
			resp := helper.HTTP.GET("http://music.163.com/api/v6/playlist/detail", params, nil)
			result := FormatList(resp.Body)
			return result
		}
	default:
		{
			return ""
		}
	}
}

func GetSong(id string, server string) string {
	switch server {
	case "netease":
		{
			params := make(map[string]any)
			params["c"] = `[{"id":`+id+`,"v":"0"}]`
			resp := helper.HTTP.GET("https://music.163.com/api/v3/song/detail/", params, nil)
			result := FormatSong(resp.Body)
			return result
		}
	default:
		{
			return ""
		}
	}
}

func GetUrl(id string, server string) string {
	switch server {
	case "netease":
		{
			params := make(map[string]any)
			params["ids"] = "["+id+"]"
			params["br"] = 320000
			resp := helper.HTTP.GET("http://music.163.com/api/song/enhance/player/url", params, nil)
			result := FormatUrl(resp.Body)
			return result
		}
	default:
		{
			return ""
		}
	}
}