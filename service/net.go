package service

import (
	"GinAPI/helper"
	"fmt"
)

func TcpPing() any {
	// m := make(map[string]string)
	resp,err := helper.Net.Tcping("119.29.94.101:80")

	fmt.Println(resp)
	fmt.Println(err)

	// if err != nil {
	// 	common.Log("app").Warn().Msg(err.Error())
	// 	return "一言获取失败"
	// }

	// fmt.Println(resp)
	return resp
}