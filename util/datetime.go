package util

import (
	"github.com/spf13/viper"
	"time"
)

func GetLocation() (*time.Location, error) {
	location := viper.GetString("timezone")
	if location == "" {
		location = "Asia/Chongqing"
	}
	loc, err := time.LoadLocation(location)

	return loc, err
}
