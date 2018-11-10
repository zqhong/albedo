package main

import (
	"github.com/zqhong/albedo/app"
	"github.com/zqhong/albedo/cmd"
	"github.com/zqhong/albedo/cronjob"
)

func main() {
	app.InitCli()

	// 注册并启动调度任务器
	cronjob.RegisterCliCron()

	cmd.Execute()
}
