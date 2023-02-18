package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func InitAPI(api *gin.RouterGroup) {
	musicAPI(api)
}

func musicAPI(route *gin.RouterGroup) {
	music := route.Group("music")
	{
		music.GET("/qq", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{"msg": "QQ音乐接口"})
		})
	}
}
