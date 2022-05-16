// @Description
// @Author wangcanjia
// @Copyright 2022 sndks.com. All rights reserved.
// @LastModify 2022/5/16 21:55

package labutils

import "testing"

func TestIsInteger(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "123",
			args: args{
				"123",
			},
			want: true,
		},
		{
			name: "123.2",
			args: args{
				"123.2",
			},
			want: false,
		},
		{
			name: "a123",
			args: args{
				"a123",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsInteger(tt.args.s); got != tt.want {
				t.Errorf("IsInteger() = %v, want %v", got, tt.want)
			}
		})
	}
}
