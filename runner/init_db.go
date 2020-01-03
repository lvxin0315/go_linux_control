package runner

import (
	"github.com/lvxin0315/go_linux_control/common"
	"github.com/lvxin0315/go_linux_control/db_conn"
	"github.com/lvxin0315/go_linux_control/model"
	"github.com/sirupsen/logrus"
	"io/ioutil"
)

//初始化数据库结构
func InitDB() {
	if fileStat, _ := common.PathExists("tmp/db_init.lock"); fileStat {
		logrus.Info("非第一次启动")
		return
	}
	//初始化db
	db, err := db_conn.GetGormDB()
	if err != nil {
		logrus.Error(err)
		panic(err)
	}
	//建表
	tableList := []interface{}{
		&model.App{},
		&model.Cmd{},
		&model.SendCmd{}}
	//先删除再创建
	db.DropTableIfExists(tableList...)
	db.CreateTable(tableList...)
	//建立lock文件
	err = ioutil.WriteFile("tmp/db_init.lock", []byte("go_linux_control"), 0666) //写入文件(字节数组)
	if err != nil {
		logrus.Error(err)
		panic(err)
	}
}
