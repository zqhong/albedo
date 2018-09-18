package router

import (
	"github.com/zqhong/albedo/app"
	"github.com/zqhong/albedo/handler/sd"
)

func RegisterApiRouter() {
	// The health check handlers
	svcd := app.Engine.Group("/sd")
	{
		svcd.GET("/health", sd.HealthCheck)
	}
}
