// @Description
// @Author wangcanjia
// @Copyright 2022 sndks.com. All rights reserved.
// @LastModify 2022/5/17 10:02

package labutils

import "testing"

func TestMd5File(t *testing.T) {
	type args struct {
		file string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "md5file.txt",
			args: args{
				"data/md5file.txt",
			},
			want:    "bf8a513a59df6b873ad06042e00cf2d2",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Md5File(tt.args.file)
			if (err != nil) != tt.wantErr {
				t.Errorf("Md5File() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Md5File() got = %v, want %v", got, tt.want)
			}
		})
	}
}
