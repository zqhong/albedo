package app

import (
	"github.com/gin-gonic/gin"
	"github.com/lexkong/log"
	"github.com/spf13/viper"
	"github.com/zqhong/albedo/router"
	"github.com/zqhong/albedo/router/middleware"
	"net/http"
)

func InitGin() {
	// Set gin mode.
	gin.SetMode(viper.GetString("runmode"))

	// Create the Gin engine.
	g := gin.New()

	// Routes.
	router.Load(
		// Cores.
		g,
		// Middlwares.
		middleware.RequestId(),
	)

	log.Debugf("Start to listening the incoming requests on http address: %s", viper.GetString("addr"))
	log.Info(http.ListenAndServe(viper.GetString("addr"), g).Error())
}
