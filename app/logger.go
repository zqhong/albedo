package app

import (
	"fmt"
	"gitee.com/zqhong/df-server/util"
	"github.com/astaxie/beego/logs"
)

func InitLogger(filePath string) {
	loggerLevel := logs.LevelInfo
	if util.IsDebugMode() {
		loggerLevel = logs.LevelDebug
	}
	loggerConf := fmt.Sprintf(`{"filename":"%s", "level": %d}`, filePath, loggerLevel)
	logs.SetLogger(logs.AdapterFile, loggerConf)

	logs.Debug("logger config: " + loggerConf)
	logs.Debug("logger init successful")
}
