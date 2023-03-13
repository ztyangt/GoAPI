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
// 	// curl, _ := ctx.GetQuery("url")
// 	agent := ctx.Query("agent")

// 	// decodedQuery, _ := url.QueryUnescape(agent)

// 	// u, _ := url.Parse()
// 	s, _ := url.QueryUnescape(ctx.Request.URL.String())
// 	surl := strings.Replace(s, " ", "a", -1)

// 	values, _ := url.ParseQuery(surl)
// 	fmt.Println("------------------------------")
// 	fmt.Println(surl)
// 	fmt.Println(agent)
// 	// fmt.Println(u.RawQuery)
// 	fmt.Println(values)
// 	// fmt.Println(u.RawQuery)
// 	// fmt.Println(ctx.Request.URL.String())

// 	// data := GetBogus(u.RawQuery, agent)
// 	ctx.JSON(http.StatusOK, "data")
// }
