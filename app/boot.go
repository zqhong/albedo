package app

import (
	"errors"
	"github.com/lexkong/log"
	"github.com/spf13/viper"
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
			log.Fatal("The router has no response, or it might took too long to start up.", err)
		}
		log.Debug("The router has been deployed successfully")
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
		log.Info("Waiting for the router, retry in 1 second.")
		time.Sleep(time.Second)
	}
	return errors.New("can not connect to the router")
}
