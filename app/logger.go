package app

import (
	"fmt"
	"github.com/astaxie/beego/logs"
	"github.com/zqhong/albedo/util"
)

func InitLogger(filePath string) {
	loggerLevel := logs.LevelInfo
	if util.IsDebugMode() {
		loggerLevel = logs.LevelDebug
	}
	loggerConf := fmt.Sprintf(`{"filename":"%s", "level": %d}`, filePath, loggerLevel)
	err := logs.SetLogger(logs.AdapterFile, loggerConf)
	if err != nil {

	}

	logs.Debug("logger config: " + loggerConf)
	logs.Debug("logger init successful")
}
