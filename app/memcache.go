package app

import (
	"github.com/astaxie/beego/cache"
	"log"
	"os"
)

var MemCache cache.Cache

func InitMemCache() {
	bm, err := cache.NewCache("memory", `{"interval":60}`)

	if err != nil {
		log.Println("创建 cache 出现问题：" + err.Error())
		os.Exit(-1)
	}

	MemCache = bm
}
