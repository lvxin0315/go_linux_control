package service

import (
	"github.com/lvxin0315/go_linux_control/db_conn"
	"github.com/lvxin0315/go_linux_control/model"
	"github.com/nats-io/go-nats"
	"github.com/sirupsen/logrus"
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
		"id": sendId,
	})
	if sendCmd.ID <= 0 {
		logrus.Error("消息记录未找到：", sendId)
		return
	}
	sendCmd.Result = string(result)
	db.Save(sendCmd)
}
