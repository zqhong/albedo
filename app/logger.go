package app

import (
	"fmt"
	"github.com/astaxie/beego/logs"
	"github.com/zqhong/albedo/util"
	"log"
	"os"
	"sync"
)

var onceLogger sync.Once

func InitLogger(filePath string) {
	onceLogger.Do(func() {
		loggerLevel := logs.LevelInfo
		if util.IsDebugMode() {
			loggerLevel = logs.LevelDebug
		}
		loggerConf := fmt.Sprintf(`{"filename":"%s", "level": %d}`, filePath, loggerLevel)
		err := logs.SetLogger(logs.AdapterFile, loggerConf)
		if err != nil {
			log.Printf("初始化 logger 服务出错：%s\n", err.Error())
			os.Exit(1)
		}

		logs.Debug("logger 配置: " + loggerConf)
		logs.Debug("logger 初始化成功")
	})
}
