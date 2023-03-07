package main

import (
	"GinAPI/common"
	"GinAPI/router"
	"GinAPI/service"
	"fmt"

	"github.com/gin-gonic/gin"
)

func init() {
	fmt.Println()
	fmt.Println("-------------------开始------------------------")
	yiyan := service.GetYiyan()
	fmt.Println(yiyan)
	fmt.Println("-------------------结束------------------------")


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
