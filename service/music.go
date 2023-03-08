package service

func GetList(id string, server string) string {
	switch server {
	case "netease":
		{

			params := make(map[string]any)
			params["n"] = 1000
			params["id"] = id
			params["s"] = 0
			params["t"] = 0
			resp := musicGET("http://music.163.com/api/v6/playlist/detail", params, "netease")
			result := FormatList(resp.Body)
			return result
		}
	default:
		{
			return "{}"
		}
	}
}

func GetSong(id string, server string) string {
	switch server {
	case "netease":
		{
			params := make(map[string]any)
			params["c"] = `[{"id":` + id + `,"v":"0"}]`
			resp := musicGET("https://music.163.com/api/v3/song/detail/", params, "netease")
			result := FormatSong(resp.Body)
			return result
		}
	default:
		{
			return "{}"
		}
	}
}

func GetUrl(id string, server string) string {
	switch server {
	case "netease":
		{
			params := make(map[string]any)
			params["ids"] = "[" + id + "]"
			params["br"] = 320000
			resp := musicGET("http://music.163.com/api/song/enhance/player/url", params, "netease")
			result := FormatUrl(resp.Body)
			return result
		}
	default:
		{
			return "{}"
		}
	}
}

func GetMV(id string, server string) string {
	switch server {
	case "netease":
		{
			params := make(map[string]any)
			params["id"] = id
			resp := musicGET("https://music.163.com/api/mv/detail", params, "netease")

			result := FormatMV(resp.Body)
			return result
		}
	default:
		{
			return "{}"
		}
	}
}

func GetComments(id string, limit string, offset string, ctype string, server string) string {
	switch server {
	case "netease":
		{
			params := make(map[string]any)
			params["limit"] = limit
			params["offset"] = offset
			resp := musicGET("http://music.163.com/api/v1/resource/comments/R_SO_4_"+id, params, "netease")
			result := FormatComments(resp.Body, ctype)
			return result
		}
	default:
		{
			return "{}"
		}
	}
}

func GetLyric(id string, server string) string {
	switch server {
	case "netease":
		{
			params := make(map[string]any)
			params["id"] = id
			params["lv"] = 1
			params["kv"] = 1
			params["tv"] = -1
			resp := musicGET("https://music.163.com/api/song/lyric", params, "netease")

			result := FormatLyric(resp.Body)
			return result
		}
	default:
		{
			return "{}"
		}
	}
}

func Search(s string, type_ string, offset string, limit string, server string) string {
	switch server {
	case "netease":
		{
			params := make(map[string]any)
			params["s"] = s
			params["type"] = type_
			params["offset"] = offset
			params["limit"] = limit
			resp := musicGET("http://music.163.com/api/cloudsearch/pc", params, "netease")
			result := FormatSearch(resp.Body, type_)
			return result
		}
	default:
		{
			return "{}"
		}
	}
}

func GetUser(id string, server string) string {
	switch server {
	case "netease":
		{
			resp := musicGET("https://music.163.com/api/v1/user/detail/"+id, nil, "netease")
			result := FormatUser(resp.Body)
			return result
		}
	default:
		{
			return "{}"
		}
	}
}