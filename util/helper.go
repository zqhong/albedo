package util

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func IsDebug() bool {
	if viper.GetString("runmode") == gin.DebugMode {
		return true
	}
	return false
}
