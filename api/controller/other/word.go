package other

import (
	"GoAPI/helper"
)

func GetYiyan() string {
	// client := resty.New()
	resp := helper.HTTP.GET("https://v1.hitokoto.cn", nil, nil)

	// fmt.Println(resp.StatusCode)
	// fmt.Println(resp.Data)
	// fmt.Println(resp.Body)
	// fmt.Println(resp.Header)
	// fmt.Println(resp.Err)

	// if err != nil {
	// 	common.Log("app").Warn().Msg(err.Error())
	// 	return "一言获取失败"
	// }

	// fmt.Println(resp)
	return resp.Body
	// resp.StatusCode()

}
