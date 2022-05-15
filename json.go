package labutils

import (
	"encoding/json"
)

func JsonToMap(data string) (map[string]any, error) {
	var dat map[string]any
	err := json.Unmarshal([]byte(data), &dat)
	return dat, err
}
func JsonToString(v any) (string, error) {
	jsonByte, err := json.Marshal(v)
	return string(jsonByte), err
}
func MapFilterToJson(v map[string]any, filterKeys ...string) (string, error) {
	if len(filterKeys) > 0 {
		dat := make(map[string]any, 0)
		for key, value := range v {
			if InSlice(key, filterKeys) {
				continue
			}
			dat[key] = value
		}
		return JsonToString(dat)
	} else {
		return JsonToString(v)
	}
}

// JsonFilter json字段过滤
// 输入：{"id":123456789012345,"name":"My Name","total":10.23}
// 输出：{"id":123456789012345,"total":10.23}
func JsonFilter(data string, filterKeys ...string) (string, error) {
	mapData, err := JsonToMap(data)
	if err != nil {
		return "", err
	}
	newData, err := MapFilterToJson(mapData, filterKeys...)
	if err != nil {
		return "", err
	}
	return newData, nil
}
