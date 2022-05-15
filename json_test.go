package labutils

import (
	"fmt"
	"testing"
)

func TestMapFilterToJson(t *testing.T) {
	data := `{"id":123456789012345,"name":"My Name","total":10.23}`
	mapData, _ := JsonToMap(data)
	fmt.Println(mapData)
	got, err := MapFilterToJson(mapData, "name")
	fmt.Println(got, err)
}

func TestJsonFilter(t *testing.T) {
	type args struct {
		data       string
		filterKeys []string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "demo1",
			args: args{
				data:       `{"id":123456789012345,"name":"My Name","total":10.23}`,
				filterKeys: nil,
			},
			want:    `{"id":123456789012345,"name":"My Name","total":10.23}`,
			wantErr: false,
		},
		{
			name: "demo2",
			args: args{
				data:       `{"id":123456789012345,"name":"My Name","total":10.23}`,
				filterKeys: []string{"name"},
			},
			want:    `{"id":123456789012345,"total":10.23}`,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := JsonFilter(tt.args.data, tt.args.filterKeys...)
			if (err != nil) != tt.wantErr {
				t.Errorf("JsonFilter() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("JsonFilter() got = %v, want %v", got, tt.want)
			}
		})
	}
}
