package leetcode

import (
	"reflect"
	"testing"
)

func Test_readBinaryWatch(t *testing.T) {
	type args struct {
		num int
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		struct {
			name string
			args args
			want []string
		}{name: "test", args: struct{ num int }{num: 1}, want: []string{
			"1:00", "2:00", "4:00", "8:00", "0:01", "0:02", "0:04", "0:08", "0:16", "0:32",
		}},
	}
	getSet := func(strs []string) map[string]struct{} {
		set := make(map[string]struct{})
		for _, s := range strs {
			set[s] = struct{}{}
		}
		return set
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := readBinaryWatch(tt.args.num)

			if !reflect.DeepEqual(getSet(got), getSet(tt.want)) {
				t.Errorf("readBinaryWatch() = %v, want %v", got, tt.want)
			}
		})
	}
}
