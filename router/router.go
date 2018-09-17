package router

import (
	"github.com/gin-gonic/gin"
	"github.com/zqhong/albedo/handler"
	"github.com/zqhong/albedo/handler/sd"
	"github.com/zqhong/albedo/pkg/errno"
)

func Load(g *gin.Engine, mw ...gin.HandlerFunc) *gin.Engine {
	g.Use(gin.Recovery())
	g.Use(mw...)

	// 404 handler
	g.NoRoute(func(c *gin.Context) {
		handler.SendResponse(c, errno.NotFound, "")
	})

	// The health check handlers
	svcd := g.Group("/sd")
	{
		svcd.GET("/health", sd.HealthCheck)
	}

	return g
}
