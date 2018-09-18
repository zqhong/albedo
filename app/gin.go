package app

import (
	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"github.com/zqhong/albedo/handler"
	"github.com/zqhong/albedo/pkg/errno"
	"github.com/zqhong/albedo/util"
)

var Engine *gin.Engine

func InitGin() {
	// Set gin mode.
	gin.SetMode(viper.GetString("runmode"))

	// Create the Gin engine.
	Engine = gin.New()

	loadRoutes(
	// Middlwares.
	)
}

func loadRoutes(mw ...gin.HandlerFunc) {
	Engine.Use(gin.Recovery())
	Engine.Use(mw...)

	// 404 handler
	Engine.NoRoute(func(c *gin.Context) {
		handler.SendResponse(c, errno.RouteNotFound, "")
	})
	Engine.NoMethod(func(c *gin.Context) {
		handler.SendResponse(c, errno.MethodNotFound, "")
	})

	if util.IsDebug() {
		// 性能分析工具
		pprof.Register(Engine)
	}
}
