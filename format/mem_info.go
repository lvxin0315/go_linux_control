package format

import (
	"encoding/json"
	"github.com/sirupsen/logrus"
	"strings"
)

//MemTotal:       16342672 kB
//MemFree:          597604 kB
//MemAvailable:   11863416 kB
//Buffers:          998572 kB
//Cached:          9609284 kB
//SwapCached:          276 kB
//Active:          5310820 kB
//Inactive:        9215580 kB
//Active(anon):    2373880 kB
//Inactive(anon):  1303736 kB
//......
type MemInfoFormat struct {
}

func (f *MemInfoFormat) ToJson(content string) string {
	//根据换行符拆分
	contents := strings.Split(content, "\n")
	if len(contents) <= 1 {
		return content
	}

	//定义map
	var items []map[string]string
	for _, value := range contents {
		item := make(map[string]string)
		kv := strings.Split(value, ":")
		if len(kv) < 2 {
			continue
		}
		item["key"] = kv[0]
		item["value"] = strings.TrimLeft(kv[1], " ")
		items = append(items, item)
	}

	jsonByte, err := json.Marshal(items)
	if err != nil {
		logrus.Error("MemInfoFormat ToJson:", err)
	}
	return string(jsonByte)
}
