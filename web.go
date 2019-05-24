package main

import (
	"fmt"
	"github.com/spf13/pflag"
	"github.com/zqhong/albedo/app"
	"github.com/zqhong/albedo/constant"
	"github.com/zqhong/albedo/router"
	"runtime"
)

var (
	webVersion    bool
	webConfigPath string
)

func init() {
	pflag.BoolVar(&webVersion, "v", false, "显示当前版本信息")
	pflag.StringVar(&webConfigPath, "c", "", "配置文件路径")
	pflag.Parse()
}

func main() {
	if webVersion {
		fmt.Printf("Albedo %s %s %s %s\n", constant.Version, runtime.GOOS, runtime.GOARCH, constant.BuildTime)
		return
	}

	app.InitConfig(webConfigPath)

	app.InitWeb()

	// 加载用户自定义的路由
	router.RegisterApiRouter()

	app.RunWeb()
}
