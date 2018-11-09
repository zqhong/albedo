package app

import (
	"fmt"
	"github.com/astaxie/beego/logs"
	"github.com/go-redis/redis"
	"github.com/spf13/viper"
	"os"
)

var Redis redis.Client

func InitRedis() {
	options := &redis.Options{
		Addr: fmt.Sprintf("%s:%d", viper.GetString("redis.host"), viper.GetInt("redis.port")),
		DB:   viper.GetInt("redis.db"),
	}
	redisPwd := viper.GetString("redis.password")
	if redisPwd != "" {
		options.Password = redisPwd
	}
	Redis = *redis.NewClient(options)

	// test
	_, err := Redis.Ping().Result()
	if err != nil {
		logs.Error(err.Error())
		os.Exit(1)
	}

}
