package router

import (
	"github.com/gin-gonic/gin"
	"github.com/zqhong/albedo/handler"
)

func RegisterApiRouter(engine *gin.Engine) {
	engine.GET("/test", func(c *gin.Context) {
		handler.SendResponse(c, nil, "test api")
	})
}
