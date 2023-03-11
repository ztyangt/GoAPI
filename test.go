package main

import (
	"GoAPI/common"
	"fmt"
	"io/ioutil"

	"github.com/dop251/goja"
)

func Test() {
	f, err := ioutil.ReadFile("source/X-Bogus.min.js")
	if err != nil {
		common.Log("app").Error().Msg("read source/X-Bogus.min.js fail")
		fmt.Println("read fail", err)
	}

	vm := goja.New()
	_, err1 := vm.RunString(string(f))
	if err1 != nil {

		fmt.Println("JS代码有问题！")
		common.Log("app").Error().Msg(err1.Error())
		return
	}
	var fn func(string, string) string
	err = vm.ExportTo(vm.Get("sign"), &fn)
	if err != nil {
		fmt.Println("Js函数映射到 Go 函数失败！")
		return
	}
	fmt.Println("斐波那契数列第30项的值为：", fn("https://www.douyin.com/aweme/v1/web/aweme/detail/?aweme_id=6772455766503755022&aid=1128&version_name=23.5.0&device_platform=android&os_version=2333", "Mozilla/5.0 (Linux; Android 8.0; Pixel 2 Build/OPD3.170816.012) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/87.0.4280.88 Mobile Safari/537.36 Edg/87.0.664.66"))
}
