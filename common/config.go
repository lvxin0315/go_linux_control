package common

import (
	"fmt"
	"github.com/Unknwon/goconfig"
	"github.com/sirupsen/logrus"
)

func GetConfig(section, key string) string {
	cfg, err := goconfig.LoadConfigFile("etc/config.ini")
	if err != nil {
		logrus.Error("GetConfig 「config.ini」 is error:", err)
		panic(err)
	}
	value, err := cfg.GetValue(section, key)
	if err != nil {
		logrus.Error(fmt.Sprintf("GetConfig 「config.ini」's %s.%s is error:", section, key), err)
		panic(err)
	}
	return value
}
