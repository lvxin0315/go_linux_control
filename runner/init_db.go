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
	//增加默认cmd内容
	db.Save(DiskUseCmd())
	db.Save(MemInfoCmd())
	db.Save(CpuInfoCmd())
	//建立lock文件
	err = ioutil.WriteFile("tmp/db_init.lock", []byte("go_linux_control"), 0666) //写入文件(字节数组)
	if err != nil {
		logrus.Error(err)
		panic(err)
	}
}

func DiskUseCmd() *model.Cmd {
	cmd := new(model.Cmd)
	cmd.Title = "查看磁盘空间"
	cmd.Cmd = "df -h"
	cmd.RouteKey = "disk_use"
	cmd.IsSystem = true
	cmd.Des = "查看磁盘空间"
	return cmd
}

func MemInfoCmd() *model.Cmd {
	cmd := new(model.Cmd)
	cmd.Title = "查看内存情况"
	cmd.Cmd = "cat /proc/meminfo"
	cmd.RouteKey = "meminfo"
	cmd.IsSystem = true
	cmd.Des = "查看内存情况"
	return cmd
}

func CpuInfoCmd() *model.Cmd {
	cmd := new(model.Cmd)
	cmd.Title = "查看cpu情况"
	cmd.Cmd = "cat /proc/cpuinfo"
	cmd.RouteKey = "cpuinfo"
	cmd.IsSystem = true
	cmd.Des = "查看cpu情况"
	return cmd
}
