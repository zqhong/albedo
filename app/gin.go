package app

import (
	"github.com/gin-gonic/gin"
	"github.com/lexkong/log"
	"github.com/spf13/viper"
	"net/http"

	"github.com/gin-contrib/pprof"
	"github.com/zqhong/albedo/handler"
	"github.com/zqhong/albedo/handler/sd"
	"github.com/zqhong/albedo/pkg/errno"
	"github.com/zqhong/albedo/router"
)

func InitGin() {
	// Set gin mode.
	gin.SetMode(viper.GetString("runmode"))

	// Create the Gin engine.
	engine := gin.New()

	loadRoutes(
		// Cores.
		engine,
		// Middlwares.
	)

	log.Debugf("Start to listening the incoming requests on http address: %s", viper.GetString("addr"))
	log.Info(http.ListenAndServe(viper.GetString("addr"), engine).Error())
}

func loadRoutes(engine *gin.Engine, mw ...gin.HandlerFunc) *gin.Engine {
	engine.Use(gin.Recovery())
	engine.Use(mw...)

	// 404 handler
	engine.NoRoute(func(c *gin.Context) {
		handler.SendResponse(c, errno.RouteNotFound, "")
	})
	engine.NoMethod(func(c *gin.Context) {
		handler.SendResponse(c, errno.MethodNotFound, "")
	})

	// The health check handlers
	svcd := engine.Group("/sd")
	{
		svcd.GET("/health", sd.HealthCheck)
	}

	if viper.GetString("runmode") == gin.DebugMode {
		// 性能分析工具
		pprof.Register(engine)
	}

	// 加载用户自定义的路由
	router.RegisterApiRouter(engine)

	return engine
}
