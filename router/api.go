package router

import (
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
		music.GET("/qq", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{"msg": "QQ音乐接口"})
		})
	}
}

func wordAPI(route *gin.RouterGroup) {
	word := route.Group("word")
	{
		word.GET("/yiyan", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, Response{200, "一言借口", "请求成功", getDeveloper()})
		})
	}
}
