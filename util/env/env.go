package env

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"github.com/zqhong/albedo/util/array"
)

func IsDebug() bool {
	if viper.GetString("runmode") == gin.DebugMode {
		return true
	}
	return false
}

func GetRunMode() string {
	runMode := viper.GetString("runMode")

	// 设置默认值
	runModeArr := []string{gin.DebugMode, gin.ReleaseMode, gin.TestMode}
	if isExists, _ := array.InArray(runMode, runModeArr); isExists == false {
		runMode = gin.ReleaseMode
	}

	return runMode
}
