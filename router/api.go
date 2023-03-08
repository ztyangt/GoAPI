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
		music.GET("/qq/:type", func(ctx *gin.Context) {

			id,_ := ctx.GetQuery("id")
			server,_ := ctx.GetQuery("server")

			type_ := ctx.Param("type")
      if type_ == "song" {
				
        data := service.GetSong(id,server)
				ctx.JSON(http.StatusOK, returnData(200, helper.Json.Decode(data),"请求成功"))

      }else if type_ == "url" {

				data := service.GetUrl(id,server)
				ctx.JSON(http.StatusOK, returnData(200, helper.Json.Decode(data),"请求成功"))
				
			}else{
				ctx.JSON(http.StatusOK, returnData(400, nil,"未知的请求方式"))
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
