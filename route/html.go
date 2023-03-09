package route

import (
	"GoAPI/api/controller/word"
	"GoAPI/common"
	"GoAPI/helper"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func Html(ctx *gin.Context) {
	year := time.Now().Year()
	yiyan := word.GetYiyan()
	data := helper.Json.Decode(yiyan).(map[string]interface{})["hitokoto"]
	ctx.HTML(http.StatusOK, "index.html", gin.H{"config": common.SiteInfo, "year": year, "yiyan": data})
}
