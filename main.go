package main

import (
	"github.com/zqhong/albedo/app"
	"github.com/zqhong/albedo/router"
)

func main() {
	app.Init()

	// 加载用户自定义的路由
	router.RegisterApiRouter()

	app.Run()
}
