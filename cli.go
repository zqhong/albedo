package main

import (
	"github.com/zqhong/albedo/app"
	"github.com/zqhong/albedo/cmd"
)

func main() {
	app.InitCli()

	cmd.Execute()
}
