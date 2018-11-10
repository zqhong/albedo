package main

import (
	"github.com/zqhong/albedo/app"
	"github.com/zqhong/albedo/cronjob"
	"github.com/zqhong/albedo/router"
)

func main() {
	app.InitWeb()

	// 加载用户自定义的路由
	router.RegisterApiRouter()

	// 注册并启动调度任务器
	cronjob.RegisterWebCron()

	app.RunWeb()
}
