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

func BogusFunc(ctx *gin.Context) {
	url, _ := ctx.GetQuery("url")
	agent, _ := ctx.GetQuery("agent")
	if url == "" || agent == "" {
		common.Log("app").Error().Msg("请求X-Bogus参数为空")
		ctx.JSON(http.StatusBadRequest, common.ReturnData(400, nil, "参数为空"))
	} else {
		data := GetBogus(url, agent)
		ctx.JSON(http.StatusOK, data)
	}

}
