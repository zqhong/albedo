package app

import (
	"fmt"
	"github.com/francoispqt/onelog"
	"github.com/zqhong/albedo/util"
	"os"
	"time"
)

var Logger *onelog.Logger

func InitLogger(filePath string) {
	f, err := os.OpenFile(filePath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println(err)
	}
	writer := os.NewFile(f.Fd(), filePath)

	if util.IsDebug() {
		Logger = onelog.New(
			writer,
			onelog.ALL,
		)
	} else {
		Logger = onelog.New(
			writer,
			onelog.INFO|onelog.WARN|onelog.ERROR|onelog.FATAL,
		)
	}
	Logger.Hook(func(e onelog.Entry) {
		e.String("time", time.Now().Format("2006/01/02 15:04:05"))
	})

	Logger.Debug("logger init successful")
}
