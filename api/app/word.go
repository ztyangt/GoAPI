package app

import (
	"GoAPI/api/controller/word"

	"github.com/gin-gonic/gin"
)

func WordAPI(route *gin.RouterGroup) {
	wordAPI := route.Group("word")
	{
		wordAPI.GET("/yiyan", word.YiyanFunc)
		// wordAPI.GET("/bogus", word.BogusFunc)
	}
}
