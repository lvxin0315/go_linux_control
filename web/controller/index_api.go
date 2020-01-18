package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/lvxin0315/go_linux_control/common"
	"github.com/lvxin0315/go_linux_control/db_conn"
	"github.com/lvxin0315/go_linux_control/model"
	"github.com/lvxin0315/go_linux_control/runner"
	"net/http"
)

func init() {
	runner.RegisterCreeperApiRunner("/api/dashboard_data", []string{"GET"}, GetIndexDatas)
}

func GetIndexDatas(c *gin.Context) {
	db := db_conn.GetGormDB()
	defer db.Close()
	op := new(Output)
	id := c.Query("id")
	if id == "" {
		op.ErrorOutput("id is nil")
		c.JSON(http.StatusOK, op)
		return
	}
	appStatus := new(model.AppStatus)
	appStatus.ID = common.StrToUint(id)

	db.First(appStatus)

	op.SuccessOutput(appStatus, "")
	c.JSON(http.StatusOK, op)
}
