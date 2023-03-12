package word

import (
	"GoAPI/common"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tidwall/gjson"
)

func YiyanFunc(ctx *gin.Context) {
	yiyan := GetYiyan()
	ctx.JSON(http.StatusOK, common.ReturnData(200, gjson.Get(yiyan, "hitokoto").String(), "请求成功"))
}

// func BogusFunc(ctx *gin.Context) {
// 	url, _ := ctx.GetQuery("url")
// 	agent, _ := ctx.GetQuery("agent")
// 	data := GetBogus(url, agent)
// 	ctx.JSON(http.StatusOK, data)
// }
