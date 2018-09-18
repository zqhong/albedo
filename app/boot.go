package app

import (
	"errors"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"net/http"
	"time"
)

func Run() {
	InitConfig()

	InitLogger()

	InitDb()

	InitGin()

	// Ping the server to make sure the router is working.
	go func() {
		if err := pingServer(); err != nil {
			Logger.Fatal("The router has no response, or it might took too long to start up",
				zap.String("err", err.Error()),
			)
		}
		Logger.Debug("The router has been deployed successfully")
	}()
}

// pingServer pings the http server to make sure the router is working.
func pingServer() error {
	for i := 0; i < 10; i++ {
		// Ping the server by sending a GET request to `/health`.
		resp, err := http.Get(viper.GetString("url") + "/sd/health")
		if err == nil && resp.StatusCode == 200 {
			return nil
		}

		// Sleep for a second to continue the next ping.
		Logger.Debug("Waiting for the router, retry in 1 second")
		time.Sleep(time.Second)
	}
	return errors.New("can not connect to the router")
}
