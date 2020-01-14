package format

import (
	"encoding/json"
	"github.com/sirupsen/logrus"
	"strings"
)

//df -h 内容处理
//	Filesystem      Size   Used  Avail Capacity iused               ifree %iused  Mounted on
//	/dev/disk2s1   223Gi  189Gi   31Gi    86% 1747322 9223372036853028485    0%   /
//	devfs          194Ki  194Ki    0Bi   100%     672                   0  100%   /dev
//	/dev/disk2s4   223Gi  2.0Gi   31Gi     6%       2 9223372036854775805    0%   /private/var/vm
//	/dev/disk0s3   476Gi  130Gi  345Gi    28%  373103           362179837    0%   /Volumes/Untitled
//	map -hosts       0Bi    0Bi    0Bi   100%       0                   0  100%   /net
//	map auto_home    0Bi    0Bi    0Bi   100%       0                   0  100%   /home
type DFHFormat struct {
}

func (f *DFHFormat) ToJson(content string) string {
	//根据换行符拆分
	contents := strings.Split(content, "\n")
	if len(contents) <= 1 {
		return content
	}
	//第一个行作为key
	keys := strings.Fields(contents[0])
	//定义map
	var items []map[string]string
	for _, value := range contents[1:] {
		item := make(map[string]string)
		fields := strings.Fields(value)
		for i, field := range fields {
			item[keys[i]] = field
		}
		items = append(items, item)
	}
	jsonByte, err := json.Marshal(items)
	if err != nil {
		logrus.Error("DFHFormat ToJson:", err)
	}
	return string(jsonByte)
}
