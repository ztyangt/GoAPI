package app

import (
	"GoAPI/api/controller/other"

	"github.com/gin-gonic/gin"
)

func SiteAPI(route *gin.RouterGroup) {
	site := route.Group("site")
	{
		site.GET("/info", other.SiteFunc)
	}
}
