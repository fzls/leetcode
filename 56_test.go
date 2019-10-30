package leetcode

import (
	"reflect"
	"testing"
)

// 2019/10/30 23:50 by fzls

func Test__merge(t *testing.T) {
	type args struct {
		intervals [][]int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{name: "test1", args: struct{ intervals [][]int }{intervals: [][]int{
			{1, 3},
			{2, 6},
			{8, 10},
			{15, 18},
		}}, want: [][]int{
			{1, 6},
			{8, 10},
			{15, 18},
		}},
		{name: "test2", args: struct{ intervals [][]int }{intervals: [][]int{
			{1, 4},
			{4, 5},
		}}, want: [][]int{
			{1, 5},
		}},
		{name: "test3", args: struct{ intervals [][]int }{intervals: [][]int{
			{1, 3},
			{4, 5},
		}}, want: [][]int{
			{1, 3},
			{4, 5},
		}},
		{name: "test4", args: struct{ intervals [][]int }{intervals: [][]int{
			{1, 4},
			{2, 3},
		}}, want: [][]int{
			{1, 4},
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := _merge(tt.args.intervals); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("_merge() = %v, want %v", got, tt.want)
			}
		})
	}
}
