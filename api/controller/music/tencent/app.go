package tencent

import (
	"GoAPI/common"
	"GoAPI/helper"

	"github.com/tidwall/gjson"
)

func GetSong(id string) common.Response {
	params := make(map[string]any)
	params["songmid"] = id
	params["platform"] = "yqq"
	params["format"] = "json"
	resp := musicGET("https://c.y.qq.com/v8/fcg-bin/fcg_play_single_song.fcg", params)
	result := FormatSong(resp.Body)
	return result
}

// QQ音乐搜索接口失效
// func Search(s string, type_ string, offset string, limit string) common.Response {
// params := make(map[string]any)
// params["w"] = s
// params["type"] = type_
// params["p"] = offset
// params["n"] = limit
// params["format"] = "json"
// resp := musicGET("https://c.y.qq.com/soso/fcgi-bin/client_search_cp", params)
// result := FormatSearch(resp.Body, type_)
// 	return "{}"
// }

func GetUrl(id string) common.Response {
	params := make(map[string]any)
	params["songmid"] = id
	params["platform"] = "yqq"
	params["format"] = "json"
	resp := musicGET("https://c.y.qq.com/v8/fcg-bin/fcg_play_single_song.fcg", params)
	result := FormatUrl(resp.Body, id)
	return result
}

func GetComments(id string, limit string, offset string, ctype string) common.Response {
	topid := gjson.Get(helper.Json.Encode(GetSong(id).Data), "id").String()
	params := make(map[string]any)
	params["format"] = "json"
	params["platform"] = "yqq.json"
	params["biztype"] = "1"
	params["topid"] = topid
	params["cmd"] = "8"
	params["pagenum"] = offset
	params["pagesize"] = limit
	resp := musicGET("https://c.y.qq.com/base/fcgi-bin/fcg_global_comment_h5.fcg", params)
	result := FormatComments(resp.Body, ctype)
	return result
}

func GetLyric(id string) any {
	params := make(map[string]any)
	params["songmid"] = id
	params["format"] = "json"
	params["nobase64"] = "1"
	params["g_tk"] = "5381"
	resp := musicGET("https://c.y.qq.com/lyric/fcgi-bin/fcg_query_lyric_new.fcg", params)
	// result := FormatLyric(resp.Body)
	return helper.Json.Decode(resp.Body)
}

func GetList(id string) common.Response {
	params := make(map[string]any)
	params["id"] = id
	params["format"] = "json"
	params["newsong"] = "1"
	params["platform"] = "jqspaframe.json"
	resp := musicGET("https://c.y.qq.com/v8/fcg-bin/fcg_v8_playlist_cp.fcg", params)
	result := FormatList(resp.Body)
	return result
}

func GetMV(id string) common.Response {
	params := make(map[string]any)
	params["g_tk"] = "5381"
	params["data"] = `{"getMvUrl":{"module":"gosrf.Stream.MvUrlProxy","method":"GetMvUrls","param":{"vids":["` + id + `"],"request_typet":10001}}}`
	resp := musicGET("https://u.y.qq.com/cgi-bin/musicu.fcg", params)
	result := FormatMV(resp.Body, id)
	return result
}

// func GetUser(id string) common.Response {
// 	resp := musicGET("https://music.163.com/api/v1/user/detail/"+id, nil)
// 	result := FormatUser(resp.Body)
// 	return result
// }
