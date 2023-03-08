package router

import (
	"GinAPI/common"
	"GinAPI/helper"
	"GinAPI/service"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func Html(ctx *gin.Context) {
	year := time.Now().Year()
	yiyan := service.GetYiyan()
	data := helper.Json.Decode(yiyan).(map[string]interface{})["hitokoto"]
	ctx.HTML(http.StatusOK, "index.html", gin.H{"config": common.SiteInfo, "year": year,"yiyan": data})
}
