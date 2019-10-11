package leetcode

import (
	"reflect"
	"testing"
)

func Test_combine(t *testing.T) {
	type args struct {
		n int
		k int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		struct {
			name string
			args args
			want [][]int
		}{name: "test", args: struct {
			n int
			k int
		}{n: 4, k: 2}, want: [][]int{
			{2, 4},
			{3, 4},
			{2, 3},
			{1, 2},
			{1, 3},
			{1, 4},
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := combine(tt.args.n, tt.args.k)
			sortIntListList(got)
			sortIntListList(tt.want)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("combine(%v) = \nres  %v, \nwant %v", tt.args, got, tt.want)
			}
		})
	}
}
