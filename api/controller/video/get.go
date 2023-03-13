package video

import (
	"GoAPI/common"
	"GoAPI/helper"
	"io/ioutil"
	"net/http"

	"github.com/dop251/goja"
	"github.com/gin-gonic/gin"
)

var Agent = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/109.0.0.0 Safari/537.36"

func DouyinGET(url string, params map[string]any) (result helper.CurlResult) {

	var headers = make(map[string]any)

	headers["referer"] = "https://www.douyin.com/"
	headers["cookie"] = "s_v_web_id=verify_leytkxgn_kvO5kOmO_SdMs_4t1o_B5ml_BUqtWM1mP6BF"
	headers["User-Agent"] = Agent

	return helper.HTTP.GET(url, params, headers)
}

func DouyinFunc(ctx *gin.Context) {
	id, _ := ctx.GetQuery("id")
	{
		data := GetDouyin(id)
		ctx.JSON(http.StatusOK, data)
	}
}

func GetBogus(url string) string {
	f, _ := ioutil.ReadFile("source/X-Bogus.min.js")

	vm := goja.New()
	_, err1 := vm.RunString(string(f))
	if err1 != nil {
		common.Log("app").Error().Msg(err1.Error())
		return "生成X-Bougus失败"
	}
	var fn func(string, string) string
	err := vm.ExportTo(vm.Get("sign"), &fn)
	if err != nil {
		common.Log("app").Error().Msg("Js函数映射到 Go 函数失败！")
		return "Js函数映射到 Go 函数失败！"
	}
	return fn(url, Agent)
}
