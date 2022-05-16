// @Description
// @Author wangcanjia
// @Copyright 2022 sndks.com. All rights reserved.
// @LastModify 2022/5/16 17:56

package labutils

import (
	"reflect"
	"strings"
	"testing"
)

func TestArrayMapForTrim(t *testing.T) {
	type args struct {
		hayStack []string
		cutset   string
		callback arrayMapForTrimCallback
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "Trim",
			args: args{
				hayStack: []string{"'DemoA'", "'DemoB'"},
				cutset:   "'",
				callback: strings.Trim,
			},
			want: []string{"DemoA", "DemoB"},
		},
		{
			name: "TrimLeft",
			args: args{
				hayStack: []string{"'DemoA'", "'DemoB'"},
				cutset:   "'",
				callback: strings.TrimLeft,
			},
			want: []string{"DemoA'", "DemoB'"},
		},
		{
			name: "TrimRight",
			args: args{
				hayStack: []string{"'DemoA'", "'DemoB'"},
				cutset:   "'",
				callback: strings.TrimRight,
			},
			want: []string{"'DemoA", "'DemoB"},
		},
		{
			name: "TrimPrefix",
			args: args{
				hayStack: []string{"'DemoA'", "'DemoB'"},
				cutset:   "'",
				callback: strings.TrimPrefix,
			},
			want: []string{"DemoA'", "DemoB'"},
		},
		{
			name: "TrimSuffix",
			args: args{
				hayStack: []string{"'DemoA'", "'DemoB'"},
				cutset:   "'",
				callback: strings.TrimSuffix,
			},
			want: []string{"'DemoA", "'DemoB"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ArrayMapForTrim(tt.args.hayStack, tt.args.cutset, tt.args.callback); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ArrayMapForTrim() = %v, want %v", got, tt.want)
			}
		})
	}
}
