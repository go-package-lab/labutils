// @Description
// @Author wangcanjia
// @Copyright 2022 sndks.com. All rights reserved.
// @LastModify 2022/5/16 20:43

package labutils

import (
	"fmt"
	"testing"
)

func TestReflectMapValueType(t *testing.T) {
	data := `{"mater_type": 1,"title": "泰坦"}`
	mapData, err := ReflectMapValueType(data)
	fmt.Println(mapData, err)
}
