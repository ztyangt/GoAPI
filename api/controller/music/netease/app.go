package netease

func GetSong(id string) string {
	params := make(map[string]any)
	params["c"] = `[{"id":` + id + `,"v":"0"}]`
	resp := musicGET("https://music.163.com/api/v3/song/detail/", params)
	result := FormatSong(resp.Body)
	return result
}

func Search(s string, type_ string, offset string, limit string) string {
	params := make(map[string]any)
	params["s"] = s
	params["type"] = type_
	params["offset"] = offset
	params["limit"] = limit
	resp := musicGET("http://music.163.com/api/cloudsearch/pc", params)
	result := FormatSearch(resp.Body, type_)
	return result
}

func GetUrl(id string) string {
	params := make(map[string]any)
	params["ids"] = "[" + id + "]"
	params["br"] = 320000
	resp := musicGET("http://music.163.com/api/song/enhance/player/url", params)
	result := FormatUrl(resp.Body)
	return result
}

func GetComments(id string, limit string, offset string, ctype string) string {
	params := make(map[string]any)
	params["limit"] = limit
	params["offset"] = offset
	resp := musicGET("http://music.163.com/api/v1/resource/comments/R_SO_4_"+id, params)
	result := FormatComments(resp.Body, ctype)
	return result
}

func GetLyric(id string) string {
	params := make(map[string]any)
	params["id"] = id
	params["lv"] = 1
	params["kv"] = 1
	params["tv"] = -1
	resp := musicGET("https://music.163.com/api/song/lyric", params)
	result := FormatLyric(resp.Body)
	return result
}

func GetList(id string) string {
	params := make(map[string]any)
	params["n"] = 1000
	params["id"] = id
	params["s"] = 0
	params["t"] = 0
	resp := musicGET("http://music.163.com/api/v6/playlist/detail", params)
	result := FormatList(resp.Body)
	return result
}

func GetMV(id string) string {
	params := make(map[string]any)
	params["id"] = id
	resp := musicGET("https://music.163.com/api/mv/detail", params)
	result := FormatMV(resp.Body, id)
	return result
}

func GetUser(id string) string {
	resp := musicGET("https://music.163.com/api/v1/user/detail/"+id, nil)
	result := FormatUser(resp.Body)
	return result
}
