package net

import (
	"GoAPI/helper"
	"fmt"
)

func TcpPing() any {
	resp, err := helper.Net.Tcping("119.29.94.101:80")

	fmt.Println(resp)
	fmt.Println(err)
	return resp
}
