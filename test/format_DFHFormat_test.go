package test

import (
	"fmt"
	"github.com/lvxin0315/go_linux_control/format"
	"testing"
)

var dfhStr = `Filesystem      Size   Used  Avail Capacity iused               ifree %iused  Mounted on
/dev/disk2s1   223Gi  189Gi   31Gi    86% 1747322 9223372036853028485    0%   /
devfs          194Ki  194Ki    0Bi   100%     672                   0  100%   /dev
/dev/disk2s4   223Gi  2.0Gi   31Gi     6%       2 9223372036854775805    0%   /private/var/vm
/dev/disk0s3   476Gi  130Gi  345Gi    28%  373103           362179837    0%   /Volumes/Untitled
map -hosts       0Bi    0Bi    0Bi   100%       0                   0  100%   /net
map auto_home    0Bi    0Bi    0Bi   100%       0                   0  100%   /home`

func Test_ToJson(t *testing.T) {
	dfhFormat := new(format.DFHFormat)
	jsonStr := dfhFormat.ToJson(dfhStr)
	fmt.Println(jsonStr)
}
