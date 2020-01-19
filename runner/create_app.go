package runner

import (
	"fmt"
	"github.com/lvxin0315/go_linux_control/common"
	"github.com/lvxin0315/go_linux_control/db_conn"
	"github.com/lvxin0315/go_linux_control/model"
	"github.com/sirupsen/logrus"
)

//创建app
func CreateApp(appName, appRemark string) {
	//生成
	db := db_conn.GetGormDB()
	defer db.Close()
	tx := db.Begin()
	defer tx.Close()
	app := new(model.App)
	app.Secret = common.UniqueId()
	app.Name = appName
	app.Remark = appRemark
	//创建应用信息
	tx.Create(app)
	if app.ID <= 0 {
		tx.Rollback()
		panic("createApp error")
	}
	//创建状态信息
	appStatus := new(model.AppStatus)
	appStatus.AppId = app.ID
	appStatus.Online = false
	tx.Create(appStatus)
	if appStatus.ID <= 0 {
		tx.Rollback()
		panic("createAppStatus error")
	}
	//成功输出 Secret
	tx.Commit()
	logrus.Println(fmt.Sprintf("createApp success, Secret is 「%s」", app.Secret))
}
