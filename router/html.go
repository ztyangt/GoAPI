package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Html(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "index.html", gin.H{})
}
