package service

import (
	"fmt"
	"github.com/lvxin0315/go_linux_control/common"
	"github.com/lvxin0315/go_linux_control/db_conn"
	"github.com/lvxin0315/go_linux_control/model"
	"github.com/nats-io/go-nats"
	"github.com/sirupsen/logrus"
	"strings"
)

func GetMessage(m *nats.Msg) {
	go SaveMessage(m.Subject, m.Data)
}

//保存返回值到db
func SaveMessage(sendId string, result []byte) {
	db, err := db_conn.GetGormDB()
	if err != nil {
		logrus.Error(err)
		panic(err)
	}

	//保存记录
	sendCmd := new(model.SendCmd)
	db.First(sendCmd, map[string]interface{}{
		"id": strings.ReplaceAll(sendId, "cmd.", ""),
	})
	if sendCmd.ID <= 0 {
		logrus.Error("消息记录未找到：", sendId)
		return
	}
	sendCmd.Result = string(result)
	db.Save(sendCmd)
}

//发布消息，先使用短连接方式发送消息
func SendCmdMessage(app *model.App, cmd *model.Cmd) {
	//生成send_cmd 记录
	sendCmd := new(model.SendCmd)
	sendCmd.AppId = app.ID
	sendCmd.CmdId = cmd.ID
	sendCmd.CmdStr = cmd.Cmd
	db, err := db_conn.GetGormDB()
	if err != nil {
		logrus.Error(err)
		return
	}
	db.Create(sendCmd)
	if sendCmd.ID <= 0 {
		logrus.Error("sendCmd保存失败")
		return
	}
	conn := common.GetConfig("nats", "host")
	nc, err := nats.Connect(conn)
	defer nc.Close()
	if err != nil {
		logrus.Error(err)
		return
	}
	logrus.Info("send cmd:", cmd.Cmd)
	//指定reply
	reply := fmt.Sprintf("cmd.%d", sendCmd.ID)
	if cmd.RouteKey != "" {
		reply = fmt.Sprintf("%s.%d", cmd.RouteKey, sendCmd.ID)
	}
	err = nc.PublishRequest(
		fmt.Sprintf("app.%s", app.Secret),
		reply,
		[]byte(cmd.Cmd))
	if err != nil {
		logrus.Error(err)
		return
	}
}
