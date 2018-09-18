package app

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"strings"
)

type Config struct {
	Name string
}

func InitConfig() error {
	cfg := pflag.StringP("config", "c", "", "config file path.")
	pflag.Parse()

	c := Config{
		Name: *cfg,
	}

	if err := c.initViper(); err != nil {
		return err
	}

	c.watchConfig()

	return nil
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

func (c *Config) watchConfig() {
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		Logger.Info(fmt.Sprintf("Config file changed: %s", e.Name))
	})
}
