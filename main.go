package main

import (
	"GinAPI/common"
	"GinAPI/router"

	"github.com/gin-gonic/gin"
)

func init() {
	// 开启日志
	common.InitLog("app", "gin")
	common.Log("app").Info().Msg("app start")
}

func main() {
	// 读取配置文件
	config := common.GetConfig()

	server := gin.Default()
	server.LoadHTMLGlob("views/*")

	apiGroup := server.Group("api")
	router.InitAPI(apiGroup)

	server.GET("/", router.Html)
	server.Run(":" + config.Site.Port)
}
