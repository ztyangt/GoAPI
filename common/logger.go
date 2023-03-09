package common

import (
	"fmt"
	"os"

	"github.com/rs/zerolog"
)

var loggers = make(map[string]zerolog.Logger)

func InitLog(names ...string) {

	_, err := os.Stat("logs")
	if err != nil {
		if os.IsNotExist(err) {
			// 创建文件夹
			os.MkdirAll("logs", os.ModePerm)
		}
	}

	for _, name := range names {
		openFile, err := os.OpenFile(fmt.Sprintf("logs/%s.log", name), os.O_CREATE|os.O_APPEND|os.O_WRONLY, os.ModeAppend|os.ModePerm)
		if err != nil {
			fmt.Printf("open log file err: %v\n", err)
			return
		}
		consoleWriter := zerolog.ConsoleWriter{Out: os.Stdout}
		multi := zerolog.MultiLevelWriter(consoleWriter, openFile)
		loggers[name] = zerolog.New(multi).With().Timestamp().Logger()
	}
}

func Log(name string) *zerolog.Logger {
	instance := loggers[name]
	return &instance
}

/*
日志等级
调用zerolog.SetGlobalLevel()设置日志等级。

panic (zerolog.PanicLevel, 5)
fatal (zerolog.FatalLevel, 4)
error (zerolog.ErrorLevel, 3)
warn (zerolog.WarnLevel, 2)
info (zerolog.InfoLevel, 1)
debug (zerolog.DebugLevel, 0)
trace (zerolog.TraceLevel, -1)

*/
