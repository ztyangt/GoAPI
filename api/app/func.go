package app

import (
	"GoAPI/api/controller/music/netease"
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
			data := netease.GetSong(id, server)
			ctx.JSON(http.StatusOK, returnData(200, helper.Json.Decode(data), "请求成功"))
		}
	case "url":
		{
			data := netease.GetUrl(id, server)
			ctx.JSON(http.StatusOK, returnData(200, helper.Json.Decode(data), "请求成功"))
		}
	case "playlist":
		{
			data := netease.GetList(id, server)
			ctx.JSON(http.StatusOK, returnData(200, helper.Json.Decode(data), "请求成功"))
		}
	case "mv":
		{
			data := netease.GetMV(id, server)
			ctx.JSON(http.StatusBadRequest, returnData(200, helper.Json.Decode(data), "请求成功"))
		}
	case "comments":
		{
			limit, _ := ctx.GetQuery("limit")
			offset, _ := ctx.GetQuery("offset")
			ctype, _ := ctx.GetQuery("ctype")
			data := netease.GetComments(id, limit, offset, ctype, server)
			ctx.JSON(http.StatusOK, returnData(200, helper.Json.Decode(data), "请求成功"))
		}
	case "lyric":
		{
			data := netease.GetLyric(id, server)
			ctx.JSON(http.StatusOK, returnData(200, helper.Json.Decode(data), "请求成功"))
		}
	case "user":
		{
			data := netease.GetUser(id, server)
			ctx.JSON(http.StatusOK, returnData(200, helper.Json.Decode(data), "请求成功"))
		}
	case "search":
		{
			s, _ := ctx.GetQuery("s")
			type_, _ := ctx.GetQuery("type")
			offset, _ := ctx.GetQuery("offset")
			limit, _ := ctx.GetQuery("limit")
			data := netease.Search(s, type_, offset, limit, server)
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
