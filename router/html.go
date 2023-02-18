package router

import (
	"GinAPI/common"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func Html(ctx *gin.Context) {
	year := time.Now().Year()
	ctx.HTML(http.StatusOK, "index.html", gin.H{"config": common.SiteInfo, "year": year})
}
