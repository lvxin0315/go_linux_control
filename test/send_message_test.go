package test

import (
	"github.com/Unknwon/goconfig"
	"github.com/lvxin0315/go_linux_control/db_conn"
	"github.com/lvxin0315/go_linux_control/model"
	"github.com/lvxin0315/go_linux_control/service"
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

func Test_sendCmd(t *testing.T) {
	app := new(model.App)
	//添加一个cmd
	cmd := new(model.Cmd)
	cmd.Cmd = "df -h"
	cmd.Title = "test cmd"
	cmd.Des = "查看磁盘使用率"
	db, _ := db_conn.GetGormDB()
	db.Create(cmd)
	db.First(app)
	//logrus.Info(app)
	//logrus.Info(cmd)
	service.SendCmdMessage(app, cmd)
}
