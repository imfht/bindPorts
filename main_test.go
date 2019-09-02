package main

import (
	"reflect"
	"testing"
)

func TestLoadPorts(t *testing.T) {
	type args struct {
		file_name string
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{name: "test port",
			args: args{
				file_name: "./testdata/ports.txt",
			},
			want: []int{80, 81}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := LoadPorts(tt.args.file_name); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LoadPorts() = %v, want %v", got, tt.want)
			}
		})
	}
}
