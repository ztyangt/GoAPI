package word

// func GetBogus(url string, agnet string) common.Response {
// 	f, err := ioutil.ReadFile("source/X-Bogus.min.js")
// 	if err != nil {
// 		common.Log("app").Error().Msg("read source/X-Bogus.min.js fail")
// 		return common.ReturnData(400, nil, "X-Bogus生成失败")
// 	}

// 	vm := goja.New()
// 	_, err1 := vm.RunString(string(f))
// 	if err1 != nil {
// 		common.Log("app").Error().Msg(err1.Error())
// 		return common.ReturnData(400, nil, err1.Error())
// 	}
// 	var fn func(string, string) string
// 	err = vm.ExportTo(vm.Get("sign"), &fn)
// 	if err != nil {
// 		common.Log("app").Error().Msg("Js函数映射到 Go 函数失败！")
// 		return common.ReturnData(400, nil, "Js函数映射到 Go 函数失败！")
// 	}
// 	return common.ReturnData(200, fn(url, agnet), "请求成功")
// }
