package app

import (
	"github.com/gin-gonic/gin"
	"github.com/thinkerou/favicon"
)

func InitServer(server *gin.Engine) {
	// server.LoadHTMLGlob("views/template/*")
	// server.StaticFS("static", http.Dir("views/assets"))
	server.Use(favicon.New("views/assets/images/favicon.svg"))

	apiGroup := server.Group("api")
	InitAPI(apiGroup)

}
