package app

import (
	"GoAPI/api/controller/music/netease"
	"GoAPI/api/controller/music/tencent"

	"github.com/gin-gonic/gin"
)

func MusicAPI(route *gin.RouterGroup) {
	musicRoute := route.Group("music")
	{
		musicRoute.GET("/:type", func(ctx *gin.Context) {
			server, _ := ctx.GetQuery("server")
			switch server {
			case "netease":
				netease.MusicFunc(ctx)
			case "tencent":
				tencent.MusicFunc(ctx)
			}
		})
	}
}
