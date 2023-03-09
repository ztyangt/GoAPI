package app

import (
	"github.com/gin-gonic/gin"
)

func InitAPI(api *gin.RouterGroup) {
	siteAPI(api)
	musicAPI(api)
	wordAPI(api)
}

func siteAPI(route *gin.RouterGroup) {
	site := route.Group("site")
	{
		site.GET("/info", SiteFunc)
	}
}

func musicAPI(route *gin.RouterGroup) {
	musicRoute := route.Group("music")
	{
		musicRoute.GET("/163/:type", MusicFunc)
	}
}

func wordAPI(route *gin.RouterGroup) {
	wordAPI := route.Group("word")
	{
		wordAPI.GET("/yiyan", YiyanFunc)
	}
}
