package service

import (
	"github.com/lvxin0315/go_linux_control/db_conn"
	"github.com/lvxin0315/go_linux_control/model"
)

func GetAppStatusList() []*model.AppStatus {
	var appStatusList []*model.AppStatus
	db := db_conn.GetGormDB()
	db.Find(&appStatusList)
	for _, appStatus := range appStatusList {
		appStatus.App = new(model.App)
		appStatus.App.ID = appStatus.AppId
		db.First(appStatus.App)
	}
	return appStatusList
}
