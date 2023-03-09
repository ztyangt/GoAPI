package main

import (
	"GoAPI/common"
	"GoAPI/route"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func init() {
	// fmt.Println()
	// fmt.Println("-------------------开始------------------------")
	// data := service.GetList("723700390","netease")
	// fmt.Println(data)
	// fmt.Println("-------------------结束------------------------")

	// 开启日志
	common.InitLog("app", "gin")
	common.Log("app").Info().Msg("app start")
}

func main() {
	// gin.SetMode(gin.ReleaseMode) //生产环境模式
	server := gin.Default()
	server.Use(cors.New(cors.Config{
		AllowAllOrigins: true, //允许所有请求源
		MaxAge:          12 * time.Hour,
	}))
	route.InitServer(server)
	server.Run(":" + common.SiteInfo.Site.Port)
}
