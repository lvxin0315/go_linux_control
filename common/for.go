package common

import (
	"github.com/sirupsen/logrus"
	"strconv"
)

func StrToInt64(str string) int64 {
	num, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		logrus.Error("StrToInt64:", err)
		panic(err)
	}
	return num
}

func StrToUint64(str string) uint64 {
	num, err := strconv.ParseUint(str, 10, 64)
	if err != nil {
		logrus.Error("StrToUint64:", err)
		panic(err)
	}
	return num
}

func StrToUint(str string) uint {
	num, err := strconv.ParseUint(str, 10, 64)
	if err != nil {
		logrus.Error("StrToUint:", err)
		panic(err)
	}
	return uint(num)
}
