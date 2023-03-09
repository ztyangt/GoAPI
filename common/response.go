package common

import (
	"time"
)

type developer struct {
	Author string `json:"author"`
	Home   string `json:"home"`
	QQ     string `json:"QQ"`
	Mail   string `json:"mail"`
	Time   string `json:"time"`
}

type Response struct {
	Code      int       `json:"code"`
	Data      any       `json:"data"`
	Msg       string    `json:"msg"`
	Developer developer `json:"developer"`
}

func GetDeveloper() developer {
	timeStr := time.Now().Format("2006-01-02 15:04:05")
	config := SiteInfo.Author
	return developer{
		Author: config.Name,
		Home:   config.Url,
		QQ:     config.QQ,
		Mail:   config.Email,
		Time:   timeStr,
	}
}

func ReturnData(code int, result any, msg string) Response {
	return Response{
		Code:      code,
		Data:      result,
		Msg:       msg,
		Developer: GetDeveloper(),
	}
}
