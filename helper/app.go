package helper

import (
	"fmt"
	"reflect"
)

func init() {
	Get.Type = getType
	In.Array = inArray
}

var Get struct {
	Type func(value any) string
}

var In struct {
	Array func(value any, array []any) bool
}

func typeof(args ...any) (typeof string, empty bool) {
	typeof, empty = "string", false
	for _, item := range args {

		// 判断是否为空
		if item == nil {
			empty = true
			continue
		}

		// 先利用反射获取数据类型，再进入不同类型的判空逻辑
		switch reflect.TypeOf(item).Kind().String() {
		case "int":
			typeof = "int"
			if item == 0 {
				empty = true
			}
		case "string":
			typeof = "string"
			if item == "" {
				empty = true
			}
		case "int64":
			typeof = "int64"
			if item == 0 {
				empty = true
			}
		case "uint8":
			typeof = "bool"
			if item == false {
				empty = true
			}
		case "float64":
			typeof = "float"
			if item == 0.0 {
				empty = true
			}
		case "byte":
			typeof = "byte"
			if item == 0 {
				empty = true
			}
		case "ptr":
			typeof = "ptr"
			// 接口状态下，它不认为自己是nil，所以要用反射判空
			if item == nil {
				empty = true
			}
			// 反射判空逻辑
			if reflect.ValueOf(item).IsNil() {
				// 利用反射直接判空
				empty = true
			}
		case "struct":
			typeof = "struct"
			if item == nil {
				empty = true
			}
		case "slice":
			typeof = "slice"
			s := reflect.ValueOf(item)
			if s.Len() == 0 {
				empty = true
			}
		case "array":
			typeof = "array"
			s := reflect.ValueOf(item)
			if s.Len() == 0 {
				empty = true
			}
		case "map":
			typeof = "map"
			s := reflect.ValueOf(item)
			if s.Len() == 0 {
				empty = true
			}
		case "chan":
			typeof = "chan"
			s := reflect.ValueOf(item)
			if s.Len() == 0 {
				empty = true
			}
		case "func":
			typeof = "func"
			if item == nil {
				empty = true
			}
		//case "bool":
		//	typeof = "bool"
		//	if item == false {
		//		empty = true
		//	}
		default:
			typeof = "other"
			empty = true
			fmt.Println("奇怪的数据类型")
		}

	}
	return
}
