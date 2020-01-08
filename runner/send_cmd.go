package runner

import (
	"github.com/lvxin0315/go_linux_control/db_conn"
	"github.com/lvxin0315/go_linux_control/model"
	"github.com/lvxin0315/go_linux_control/service"
	"github.com/sirupsen/logrus"
)

func SendCmd(appId, cmdId string) {
	logrus.Info("send cmd ", appId, "  ", cmdId)
	db, err := db_conn.GetGormDB()
	if err != nil {
		logrus.Error(err)
		panic(err)
	}

	app := new(model.App)
	db.First(app, map[string]interface{}{
		"id": appId,
	})
	cmd := new(model.Cmd)
	db.First(cmd, map[string]interface{}{
		"id": cmdId,
	})
	logrus.Info(cmd)
	service.SendCmdMessage(app, cmd)
}
