package runner

import (
	"fmt"
	"github.com/lvxin0315/go_linux_control/db_conn"
	"github.com/lvxin0315/go_linux_control/model"
	"github.com/sirupsen/logrus"
)

//创建cmd
func CreateCmd(cmdTitle, cmdStr, cmdDes string) {
	//生成
	db, err := db_conn.GetGormDB()
	if err != nil {
		logrus.Error(err)
		panic(err)
	}
	cmd := new(model.Cmd)
	cmd.Title = cmdTitle
	cmd.Des = cmdDes
	cmd.Cmd = cmdStr
	db.Create(cmd)
	if cmd.ID <= 0 {
		panic("CreateCmd error")
	}
	//成功输出 ID
	logrus.Println(fmt.Sprintf("CreateCmd success, Id is 「%d」", cmd.ID))
}
