package cronjob

import (
	"github.com/astaxie/beego/logs"
	"github.com/robfig/cron"
)

func RegisterCliCron() {
	c := cron.New()

	c.AddFunc("@hourly", CronCliExample)

	c.Start()
	logs.Debug("cron init successful")
}

func RegisterWebCron() {
	c := cron.New()

	c.AddFunc("@hourly", CronWebExample)

	c.Start()
	logs.Debug("cron init successful")
}
