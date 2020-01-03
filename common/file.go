package common

import (
	"fmt"
	"io/ioutil"
	"os"
	"time"
)

func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func SaveShellFile(cmdStr []byte) (shellFileName string, err error) {
	shellFileName = fmt.Sprintf("/tmp/%d.sh", time.Now().UnixNano())
	err = ioutil.WriteFile(shellFileName, cmdStr, os.ModePerm)
	return shellFileName, err
}

func DeleteShellFile(shellFileName string) {
	_ = os.Remove(shellFileName)
}
