package sd

import (
	"github.com/gin-gonic/gin"
	"github.com/zqhong/albedo/handler"
)

// HealthCheck shows `OK` as the ping-pong result.
func HealthCheck(c *gin.Context) {
	handler.SendResponse(c, nil, "OK")
}
