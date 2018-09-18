package sd

import (
	"github.com/gin-gonic/gin"
	"github.com/zqhong/albedo/app"
	"github.com/zqhong/albedo/handler"
	"go.uber.org/zap"
)

// HealthCheck shows `OK` as the ping-pong result.
func HealthCheck(c *gin.Context) {
	app.Logger.Info("zap logger",
		zap.String("url", "http://www.qq.com"),
		zap.Int("num", 3),
	)

	handler.SendResponse(c, nil, "OK")
}
