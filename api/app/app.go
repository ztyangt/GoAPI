package app

import (
	"github.com/gin-gonic/gin"
)

func InitAPI(api *gin.RouterGroup) {
	SiteAPI(api)
	MusicAPI(api)
	VideoAPI(api)
	WordAPI(api)
}
