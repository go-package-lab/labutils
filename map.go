package labutils

import (
	"encoding/json"
	"reflect"
)

// ReflectMapValueType 反射map字段值类型
// Example:
// input: {"id": 1,"title": "泰坦"}
// output: map[id:float64 title:string]
func ReflectMapValueType(jsonStr string) (map[string]string, error) {
	fieldType := make(map[string]string, 0)
	mData := map[string]any{}
	err := json.Unmarshal([]byte(jsonStr), &mData)
	if err != nil {
		return fieldType, err
	}
	for k, d := range mData {
		fieldType[k] = reflect.TypeOf(d).String()
	}
	return fieldType, nil
}
