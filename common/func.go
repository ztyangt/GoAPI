package common

import (
	"github.com/tidwall/gjson"
)

type KeyValue struct {
	Key   string
	Value any
}

func ToMap(params string) map[string]any {

	result := make(map[string]any)
	data := gjson.Parse(params)

	data.ForEach(func(key, value gjson.Result) bool {
		result[key.String()] = value.String()
		return true
	})
	return result
}
