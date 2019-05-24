package main

import "C"
import (
	"fmt"
	"github.com/spf13/pflag"
	"github.com/zqhong/albedo/app"
	"github.com/zqhong/albedo/cmd"
	"github.com/zqhong/albedo/constant"
	"runtime"
)

var (
	cliVersion    bool
	cliConfigPath string
)

func init() {
	pflag.BoolVar(&cliVersion, "v", false, "显示当前版本信息")
	pflag.StringVar(&cliConfigPath, "c", "", "配置文件路径")
	pflag.Parse()
}

func main() {
	if cliVersion {
		fmt.Printf("Albedo %s %s %s %s\n", constant.Version, runtime.GOOS, runtime.GOARCH, constant.BuildTime)
		return
	}

	app.InitConfig(cliConfigPath)

	app.InitCli()

	cmd.Execute()
}
