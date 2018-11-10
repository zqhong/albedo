package cronjob

import (
	"fmt"
	"github.com/astaxie/beego/logs"
	"github.com/zqhong/albedo/app"
)

func CronCliExample() {
	fmt.Println(app.Redis)
	logs.Debug("Every hour - cli")
}
