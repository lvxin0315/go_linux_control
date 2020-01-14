package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/lvxin0315/go_linux_control/db_conn"
	"github.com/lvxin0315/go_linux_control/model"
	"github.com/lvxin0315/go_linux_control/runner"
	"net/http"
)

type DashboardData struct {
	AppList     []*model.App
	CmdList     []*model.Cmd
	SendCmdList []*model.SendCmd
}

func init() {
	runner.RegisterCreeperApiRunner("/dashboard", []string{"POST", "GET"}, Dashboard)
}

func Dashboard(c *gin.Context) {
	appList, cmdList, sendCmdList := getDatas()
	c.HTML(http.StatusOK, "dashboard.html", DashboardData{
		appList,
		cmdList,
		sendCmdList,
	})
}

func getDatas() (appList []*model.App, cmdList []*model.Cmd, sendCmdList []*model.SendCmd) {
	db, err := db_conn.GetGormDB()
	if err != nil {
		return nil, nil, nil
	}
	db.Find(&appList)
	db.Find(&cmdList)
	db.Preload("App").Preload("Cmd").Find(&sendCmdList)

	return appList, cmdList, sendCmdList
}
