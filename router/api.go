package router

import (
	"GinAPI/helper"
	"GinAPI/service"
	"net/http"

	"github.com/gin-gonic/gin"
)


func InitAPI(api *gin.RouterGroup) {
	musicAPI(api)
	wordAPI(api)
}


func musicAPI(route *gin.RouterGroup) {
	music := route.Group("music")
	{
		music.GET("/163/:type", func(ctx *gin.Context) {
			id,_ := ctx.GetQuery("id")
			server,_ := ctx.GetQuery("server")
			type_ := ctx.Param("type")
			switch type_ {

      case "song": {
        data := service.GetSong(id,server)
				ctx.JSON(http.StatusOK, returnData(200, helper.Json.Decode(data),"请求成功"))
      }
			case "url": {
				data := service.GetUrl(id,server)
				ctx.JSON(http.StatusOK, returnData(200, helper.Json.Decode(data),"请求成功"))
      }
			case "playlist": {
				data := service.GetList(id,server)
				ctx.JSON(http.StatusOK, returnData(200, helper.Json.Decode(data),"请求成功"))
      }
			case "mv": {
				data := service.GetMV(id,server)
				ctx.JSON(http.StatusBadRequest, returnData(200, helper.Json.Decode(data),"请求成功"))
      }
			case "comments": {
				limit,_ := ctx.GetQuery("limit")
				offset,_ := ctx.GetQuery("offset")
				ctype,_ := ctx.GetQuery("ctype")
				data := service.GetComments(id,limit,offset,ctype,server)
				ctx.JSON(http.StatusOK, returnData(200, helper.Json.Decode(data),"请求成功"))
			}
			case "lyric": {
				data := service.GetLyric(id,server)
				ctx.JSON(http.StatusOK, returnData(200, helper.Json.Decode(data),"请求成功"))
			}
			case "user": {
				data := service.GetUser(id,server)
				ctx.JSON(http.StatusOK, returnData(200, helper.Json.Decode(data),"请求成功"))
			}
			case "search": {
				s,_ := ctx.GetQuery("s")
				type_,_ := ctx.GetQuery("type")
				offset,_ := ctx.GetQuery("offset")
				limit,_ := ctx.GetQuery("limit")
				data := service.Search(s, type_,offset,limit,server)
				ctx.JSON(http.StatusOK, returnData(200, helper.Json.Decode(data),"请求成功"))
			}
			default: {
				ctx.JSON(http.StatusOK, returnData(400, nil,"未知的请求方式"))
			}
		}
		})
	}
}

func wordAPI(route *gin.RouterGroup) {
	word := route.Group("word")
	{
		word.GET("/yiyan", func(ctx *gin.Context) {
			yiyan := service.GetYiyan()
			ctx.JSON(http.StatusOK, Response{200, helper.Json.Decode(yiyan).(map[string]interface{})["hitokoto"], "请求成功", getDeveloper()})
		})
	}
}
