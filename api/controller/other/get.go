package other

import (
	"GoAPI/common"
	"GoAPI/helper"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SiteFunc(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, common.ReturnData(200, helper.Json.Decode(helper.Json.Encode(common.SiteInfo)), "请求成功"))
}
