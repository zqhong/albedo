package app

import (
	"github.com/astaxie/beego/logs"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"log"
	"os"
	"strings"
)

type Config struct {
	Name string
}

func InitConfig() {
	cfg := pflag.StringP("config", "c", "", "config file path.")
	pflag.Parse()

	c := Config{
		Name: *cfg,
	}

	if err := c.initViper(); err != nil {
		log.Printf("初始化 config 服务出错：%s\n", err.Error())
		os.Exit(1)
	}

	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		logs.Debug("Config file changed:" + e.Name)
	})
}

func (c *Config) initViper() error {
	if c.Name != "" {
		viper.SetConfigFile(c.Name)
	} else {
		viper.AddConfigPath("conf")
		viper.SetConfigName("config")
	}
	viper.SetConfigType("yaml")
	viper.AutomaticEnv()
	viper.SetEnvPrefix("ALBEDO")
	replacer := strings.NewReplacer(".", "_")
	viper.SetEnvKeyReplacer(replacer)

	if err := viper.ReadInConfig(); err != nil {
		return err
	}

	return nil
}
