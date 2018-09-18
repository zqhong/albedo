package router

import (
	"github.com/gin-gonic/gin"
	"github.com/zqhong/albedo/handler/sd"
)

func RegisterApiRouter(engine *gin.Engine) {
	// The health check handlers
	svcd := engine.Group("/sd")
	{
		svcd.GET("/health", sd.HealthCheck)
	}
}
