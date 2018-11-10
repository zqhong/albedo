package app

import (
	"github.com/astaxie/beego/logs"
	"github.com/robfig/cron"
)

// 参考文档：https://godoc.org/github.com/robfig/cron
func InitCronCli() {
	c := cron.New()

	c.AddFunc("@hourly", func() {
		logs.Debug("Every hour - cli")
	})

	c.Start()
	logs.Debug("cron init successful")
}

func InitCronWeb() {
	c := cron.New()

	c.AddFunc("@hourly", func() {
		logs.Debug("Every hour - web")
	})

	c.Start()
	logs.Debug("cron init successful")
}
