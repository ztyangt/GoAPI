package app

import (
	"GoAPI/api/controller/music/netease"

	"github.com/gin-gonic/gin"
)

func MusicAPI(route *gin.RouterGroup) {
	musicRoute := route.Group("music")
	{
		musicRoute.GET("/163/:type", netease.MusicFunc)
	}
}
