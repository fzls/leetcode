package leetcode

import (
	"reflect"
	"sort"
	"testing"
)

// 2019/10/04 20:47 by fzls

func Test_combinationSum(t *testing.T) {
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
		}{candidates: []int{2, 3, 6, 7}, target: 7}, want: [][]int{
			{7},
			{2, 2, 3},
		}},
		{name: "test", args: struct {
			candidates []int
			target     int
		}{candidates: []int{2, 3, 5}, target: 8}, want: [][]int{
			{2, 2, 2, 2},
			{2, 3, 3},
			{3, 5},
		}},
		{name: "test", args: struct {
			candidates []int
			target     int
		}{candidates: []int{7, 3, 2}, target: 18}, want: [][]int{
			{2, 2, 2, 2, 2, 2, 2, 2, 2},
			{2, 2, 2, 2, 2, 2, 3, 3},
			{2, 2, 2, 2, 3, 7},
			{2, 2, 2, 3, 3, 3, 3},
			{2, 2, 7, 7},
			{2, 3, 3, 3, 7},
			{3, 3, 3, 3, 3, 3},
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := combinationSum(tt.args.candidates, tt.args.target)

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
				t.Errorf("combinationSum(%v) \nres  %v, \nwant %v", tt.args, got, tt.want)
			}
		})
	}
}
