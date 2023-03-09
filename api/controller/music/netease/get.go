package netease

import (
	"GoAPI/common"
	"GoAPI/helper"
	"net/http"

	"github.com/gin-gonic/gin"
)

func musicGET(url string, params map[string]any) (result helper.CurlResult) {

	var headers = make(map[string]any)

	headers["Referer"] = "https://music.163.com/"
	headers["cookie"] = "appver=8.2.30; os=iPhone OS; osver=15.0; EVNSM=1.0.0; buildver=2206; channel=distribution; machineid=iPhone13.3"
	headers["User-Agent"] = "Mozilla/5.0 (iPhone; CPU iPhone OS 15_0 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Mobile/15E148 CloudMusic/0.1.1 NeteaseMusic/8.2.30"

	return helper.HTTP.GET(url, params, headers)
}

func MusicFunc(ctx *gin.Context) {
	id, _ := ctx.GetQuery("id")
	server, _ := ctx.GetQuery("server")
	type_ := ctx.Param("type")
	switch type_ {

	case "song":
		{
			data := GetSong(id, server)
			ctx.JSON(http.StatusOK, common.ReturnData(200, helper.Json.Decode(data), "请求成功"))
		}
	case "url":
		{
			data := GetUrl(id, server)
			ctx.JSON(http.StatusOK, common.ReturnData(200, helper.Json.Decode(data), "请求成功"))
		}
	case "playlist":
		{
			data := GetList(id, server)
			ctx.JSON(http.StatusOK, common.ReturnData(200, helper.Json.Decode(data), "请求成功"))
		}
	case "mv":
		{
			data := GetMV(id, server)
			ctx.JSON(http.StatusBadRequest, common.ReturnData(200, helper.Json.Decode(data), "请求成功"))
		}
	case "comments":
		{
			limit, _ := ctx.GetQuery("limit")
			offset, _ := ctx.GetQuery("offset")
			ctype, _ := ctx.GetQuery("ctype")
			data := GetComments(id, limit, offset, ctype, server)
			ctx.JSON(http.StatusOK, common.ReturnData(200, helper.Json.Decode(data), "请求成功"))
		}
	case "lyric":
		{
			data := GetLyric(id, server)
			ctx.JSON(http.StatusOK, common.ReturnData(200, helper.Json.Decode(data), "请求成功"))
		}
	case "user":
		{
			data := GetUser(id, server)
			ctx.JSON(http.StatusOK, common.ReturnData(200, helper.Json.Decode(data), "请求成功"))
		}
	case "search":
		{
			s, _ := ctx.GetQuery("s")
			type_, _ := ctx.GetQuery("type")
			offset, _ := ctx.GetQuery("offset")
			limit, _ := ctx.GetQuery("limit")
			data := Search(s, type_, offset, limit, server)
			ctx.JSON(http.StatusOK, common.ReturnData(200, helper.Json.Decode(data), "请求成功"))
		}
	default:
		{
			ctx.JSON(http.StatusOK, common.ReturnData(400, nil, "未知的请求方式"))
		}
	}
}
