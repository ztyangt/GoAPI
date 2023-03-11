package app

import (
	"GoAPI/api/controller/video"

	"github.com/gin-gonic/gin"
)

func VideoAPI(route *gin.RouterGroup) {
	videoRoute := route.Group("video")
	{
		videoRoute.GET("/:type", func(ctx *gin.Context) {
			type_ := ctx.Param("type")
			switch type_ {
			case "douyin":
				video.DouyinFunc(ctx)
			}
		})
	}
}
