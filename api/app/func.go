package app

import (
	"GoAPI/api/controller/music"
	"GoAPI/api/controller/word"
	"GoAPI/common"
	"GoAPI/helper"
	"net/http"

	"github.com/gin-gonic/gin"
)

func MusicFunc(ctx *gin.Context) {
	id, _ := ctx.GetQuery("id")
	server, _ := ctx.GetQuery("server")
	type_ := ctx.Param("type")
	switch type_ {

	case "song":
		{
			data := music.GetSong(id, server)
			ctx.JSON(http.StatusOK, returnData(200, helper.Json.Decode(data), "请求成功"))
		}
	case "url":
		{
			data := music.GetUrl(id, server)
			ctx.JSON(http.StatusOK, returnData(200, helper.Json.Decode(data), "请求成功"))
		}
	case "playlist":
		{
			data := music.GetList(id, server)
			ctx.JSON(http.StatusOK, returnData(200, helper.Json.Decode(data), "请求成功"))
		}
	case "mv":
		{
			data := music.GetMV(id, server)
			ctx.JSON(http.StatusBadRequest, returnData(200, helper.Json.Decode(data), "请求成功"))
		}
	case "comments":
		{
			limit, _ := ctx.GetQuery("limit")
			offset, _ := ctx.GetQuery("offset")
			ctype, _ := ctx.GetQuery("ctype")
			data := music.GetComments(id, limit, offset, ctype, server)
			ctx.JSON(http.StatusOK, returnData(200, helper.Json.Decode(data), "请求成功"))
		}
	case "lyric":
		{
			data := music.GetLyric(id, server)
			ctx.JSON(http.StatusOK, returnData(200, helper.Json.Decode(data), "请求成功"))
		}
	case "user":
		{
			data := music.GetUser(id, server)
			ctx.JSON(http.StatusOK, returnData(200, helper.Json.Decode(data), "请求成功"))
		}
	case "search":
		{
			s, _ := ctx.GetQuery("s")
			type_, _ := ctx.GetQuery("type")
			offset, _ := ctx.GetQuery("offset")
			limit, _ := ctx.GetQuery("limit")
			data := music.Search(s, type_, offset, limit, server)
			ctx.JSON(http.StatusOK, returnData(200, helper.Json.Decode(data), "请求成功"))
		}
	default:
		{
			ctx.JSON(http.StatusOK, returnData(400, nil, "未知的请求方式"))
		}
	}
}

func SiteFunc(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, returnData(200, helper.Json.Decode(helper.Json.Encode(common.SiteInfo)), "请求成功"))
}

func YiyanFunc(ctx *gin.Context) {
	yiyan := word.GetYiyan()
	ctx.JSON(http.StatusOK, Response{200, helper.Json.Decode(yiyan).(map[string]interface{})["hitokoto"], "请求成功", getDeveloper()})
}
