package word

import (
	"GoAPI/helper"
)

func GetYiyan() string {
	resp := helper.HTTP.GET("https://v1.hitokoto.cn", nil, nil)
	return resp.Body
}
