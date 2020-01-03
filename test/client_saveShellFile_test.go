package test

import (
	"github.com/lvxin0315/go_linux_control/common"
	"log"
	"testing"
)

func Test_SaveShellFile(t *testing.T) {
	shFile, err := common.SaveShellFile([]byte("df -h /"))
	defer common.DeleteShellFile(shFile)
	if err != nil {
		t.Fail()
		panic(err)
	}
	log.Println(shFile)
}
