package leetcode

import (
	"reflect"
	"sort"
	"testing"
)

// 2019/10/07 23:07 by fzls

func Test_combinationSum2(t *testing.T) {
	type args struct {
		candidates []int
		target     int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{name: "test", args: struct {
			candidates []int
			target     int
		}{candidates: []int{10, 1, 2, 7, 6, 1, 5}, target: 8}, want: [][]int{
			{1, 7},
			{1, 2, 5},
			{2, 6},
			{1, 1, 6},
		}},
		{name: "test", args: struct {
			candidates []int
			target     int
		}{candidates: []int{2, 5, 2, 1, 2}, target: 5}, want: [][]int{
			{1, 2, 2},
			{5},
		}},
		{name: "test", args: struct {
			candidates []int
			target     int
		}{candidates: []int{3, 1, 3, 5, 1, 1}, target: 8}, want: [][]int{
			{1, 1, 1, 5},
			{1, 1, 3, 3},
			{3, 5},
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := combinationSum2(tt.args.candidates, tt.args.target)

			var less func(a, b []int) bool
			less = func(a, b []int) bool {
				if len(a) == 0 && len(b) == 0 {
					return true
				} else if len(a) == 0 {
					return true
				} else if len(b) == 0 {
					return false
				} else {
					if a[0] != b[0] {
						return a[0] < b[0]
					}
					return less(a[1:], b[1:])
				}
			}

			sortFunc := func(data [][]int) {
				for i := 0; i < len(data); i++ {
					sort.Ints(data[i])
				}
				sort.Slice(data, func(i, j int) bool {
					return less(data[i], data[j])
				})
			}
			sortFunc(got)
			sortFunc(tt.want)

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("combinationSum2(%v) \nres  %v, \nwant %v", tt.args, got, tt.want)
			}
		})
	}
}
