package tencent

import (
	"GoAPI/common"
	"GoAPI/helper"
	"net/http"

	"github.com/gin-gonic/gin"
)

func musicGET(url string, params map[string]any) (result helper.CurlResult) {

	var headers = make(map[string]any)

	headers["Referer"] = "http://y.qq.com"
	headers["cookie"] = "pgv_pvi=22038528; pgv_si=s3156287488; pgv_pvid=5535248600; yplayer_open=1; ts_last=y.qq.com/portal/player.html; ts_uid=4847550686; yq_index=0; qqmusic_fromtag=66; player_exist=1"
	headers["User-Agent"] = "QQ%E9%9F%B3%E4%B9%90/54409 CFNetwork/901.1 Darwin/17.6.0 (x86_64)"
	headers["Host"] = "c.y.qq.com"

	return helper.HTTP.GET(url, params, headers)
}

func MusicFunc(ctx *gin.Context) {
	id, _ := ctx.GetQuery("id")
	type_ := ctx.Param("type")
	switch type_ {

	case "song":
		{
			data := GetSong(id)
			ctx.JSON(http.StatusOK, data)
		}
	case "url":
		{
			data := GetUrl(id)
			ctx.JSON(http.StatusOK, data)
		}

	case "comments":
		{
			limit, _ := ctx.GetQuery("limit")
			offset, _ := ctx.GetQuery("offset")
			ctype, _ := ctx.GetQuery("ctype")
			data := GetComments(id, limit, offset, ctype)
			ctx.JSON(http.StatusOK, data)
		}

	case "lyric":
		{
			data := GetLyric(id)
			ctx.JSON(http.StatusOK, data)
		}

	case "playlist":
		{
			data := GetList(id)
			ctx.JSON(http.StatusOK, data)
		}
	case "mv":
		{
			data := GetMV(id)
			ctx.JSON(http.StatusBadRequest, data)
		}

	// case "user":
	// 	{
	// 		data := GetUser(id)
	// 		ctx.JSON(http.StatusOK, data)
	// 	}
	default:
		{
			ctx.JSON(http.StatusOK, common.ReturnData(400, nil, "未知的请求方式"))
		}
	}
}
