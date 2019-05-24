package app

import (
	"fmt"
	"github.com/astaxie/beego/cache"
	"log"
	"os"
	"sync"
	"time"
)

var (
	onceMemCache sync.Once
	MemCache     cache.Cache
)

func InitMemCache() {
	onceMemCache.Do(func() {
		bm, err := cache.NewCache("memory", `{"interval":60}`)

		if err != nil {
			log.Println("创建 cache 出现问题：" + err.Error())
			os.Exit(-1)
		}

		MemCache = bm

		// test
		testKey := fmt.Sprintf("test-%s", time.Now().Format(time.StampNano))
		err = MemCache.Put(testKey, 1, 1*time.Microsecond)
		if err != nil {
			log.Println("MemCache.Put 出现问题：" + err.Error())
			os.Exit(-1)
		}
	})
}
