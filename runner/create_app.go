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
	db, err := db_conn.GetGormDB()
	if err != nil {
		logrus.Error(err)
		panic(err)
	}
	app := new(model.App)
	app.Secret = common.UniqueId()
	app.Name = appName
	app.Remark = appRemark
	db.Create(app)
	if app.ID <= 0 {
		panic("createApp error")
	}
	//成功输出 Secret
	logrus.Println(fmt.Sprintf("createApp success, Secret is 「%s」", app.Secret))
}
