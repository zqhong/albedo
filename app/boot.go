package app

import (
	"errors"
	"fmt"
	"github.com/astaxie/beego/logs"
	"github.com/fvbock/endless"
	"github.com/spf13/viper"
	"github.com/zqhong/albedo/constant"
	"github.com/zqhong/albedo/util"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

func InitWeb() {
	InitConfig()

	InitPath()

	InitLogger(constant.Path.RootDir + "/runtime/log/albedo-web.log")

	InitEnv()

	InitDb()

	InitRedis()

	InitMemCache()

	InitGin()
}

func InitCli() {
	InitConfig()

	InitPath()

	InitLogger(constant.Path.RootDir + "/runtime/log/albedo-cli.log")

	InitEnv()

	InitDb()

	InitRedis()

	InitMemCache()
}

func InitPath() {
	if filepath.IsAbs(Config.Name) == false {
		currentDir, _ := os.Getwd()
		configFilePath := filepath.Join(currentDir, Config.Name)
		configPath := filepath.Dir(configFilePath)

		constant.SetConfDir(configPath)
		constant.SetRootDir(filepath.Dir(constant.Path.ConfDir))
	} else {
		constant.SetConfDir(filepath.Dir(Config.Name))
		constant.SetRootDir(filepath.Dir(constant.Path.ConfDir))
	}
}

func InitEnv() {
	loc, err := util.GetLocation()

	if err != nil {
		log.Println("时区设置失败：" + err.Error())
		os.Exit(1)
	}

	time.Local = loc
	logs.Debug(fmt.Sprintf("默认时区（%s）设置成功", loc.String()))
}

func RunWeb() {
	logs.Debug(fmt.Sprintf("Start to listening the incoming requests on http address: %s", viper.GetString("addr")))
	logs.Info(endless.ListenAndServe(viper.GetString("addr"), Engine).Error())

	checkServer()
}

func checkServer() {
	// Ping the server to make sure the router is working.
	go func() {
		if err := pingServer(); err != nil {
			logs.Error(fmt.Sprintf("The router has no response, or it might took too long to start up, err: %s", err.Error()))
		}
		logs.Debug("The router has been deployed successfully")
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
		logs.Debug("Waiting for the router, retry in 1 second")
		time.Sleep(time.Second)
	}
	return errors.New("can not connect to the router")
}
