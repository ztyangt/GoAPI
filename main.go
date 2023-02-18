package main

import (
	"GinAPI/common"
	"GinAPI/router"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/thinkerou/favicon"
)

func init() {
	// 开启日志
	common.InitLog("app", "gin")
	common.Log("app").Info().Msg("app start")
}

func main() {
	server := gin.Default()
	server.LoadHTMLGlob("views/template/*")
	server.StaticFS("static", http.Dir("views/assets"))
	server.Use(favicon.New("./views/assets/images/favicon.ico"))

	apiGroup := server.Group("api")
	router.InitAPI(apiGroup)

	server.GET("/", router.Html)
	server.Run(":" + common.SiteInfo.Site.Port)
}
