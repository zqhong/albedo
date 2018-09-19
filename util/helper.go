package util

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"path"
	"runtime"
)

func IsDebug() bool {
	if viper.GetString("runmode") == gin.DebugMode {
		return true
	}
	return false
}

func GetCurrentFile() string {
	_, filename, _, _ := runtime.Caller(1)
	return filename
}

func GetCurrentDir() string {
	_, filename, _, _ := runtime.Caller(1)
	return path.Dir(filename)
}

func GetRootDir() string {
	currentDir := GetCurrentDir()
	return path.Dir(currentDir)
}
