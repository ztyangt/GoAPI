package video

import (
	"GoAPI/helper"
	"net/http"

	"github.com/gin-gonic/gin"
)

func DouyinGET(url string, params map[string]any) (result helper.CurlResult) {

	var headers = make(map[string]any)

	headers["referer"] = "https://www.douyin.com/"
	headers["cookie"] = "s_v_web_id=verify_leytkxgn_kvO5kOmO_SdMs_4t1o_B5ml_BUqtWM1mP6BF"
	headers["User-Agent"] = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/109.0.0.0 Safari/537.36"

	return helper.HTTP.GET(url, params, headers)
}

func DouyinFunc(ctx *gin.Context) {
	id, _ := ctx.GetQuery("id")
	{
		data := GetDouyin(id)
		ctx.JSON(http.StatusOK, data)
	}
}
