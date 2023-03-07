package main

import (
	"GinAPI/common"
	"GinAPI/router"

	"github.com/gin-gonic/gin"
)

func init() {
	common.InitSqlite()
	// 开启日志
	common.InitLog("app", "gin")
	common.Log("app").Info().Msg("app start")
}

func main() {
	server := gin.Default()
	router.InitServer(server)
	server.Run(":" + common.SiteInfo.Site.Port)
}
