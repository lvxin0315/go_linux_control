package test

import (
	"github.com/Unknwon/goconfig"
	"github.com/nats-io/go-nats"
	"github.com/sirupsen/logrus"
	"testing"
)

func Test_send(t *testing.T) {
	cfg, err := goconfig.LoadConfigFile("etc/config.ini")
	if err != nil {
		logrus.Error("goconfig.LoadConfigFile 「config.ini」 is error:", err)
		panic(err)
	}
	conn, err := cfg.GetValue("nats", "host")
	if err != nil {
		logrus.Error("cfg.GetValue 「config.ini」 is error:", err)
		panic(err)
	}
	nc, err := nats.Connect(conn)
	defer nc.Close()
	if err != nil {
		panic(err)
	}
	nc.Publish("sss", []byte("123123123"))
	nc.Publish("sss", []byte("123123123"))
	nc.Publish("sss", []byte("123123123"))
}
