package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/lvxin0315/go_linux_control/model"
	"github.com/lvxin0315/go_linux_control/runner"
	"github.com/lvxin0315/go_linux_control/service"
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
	//获取应用列表
	appStatusList := service.GetAppStatusList()
	op := new(Output)
	op.Data = appStatusList
	c.HTML(http.StatusOK, "dashboard.html", op)
}
