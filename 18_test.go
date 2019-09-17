package leetcode

import (
	"reflect"
	"sort"
	"testing"
)

func Test_fourSum(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{name: "test", args: struct {
			nums   []int
			target int
		}{nums: []int{1, 0, -1, 0, -2, 2}, target: 0}, want: [][]int{
			[]int{-1, 0, 0, 1},
			[]int{-2, -1, 1, 2},
			[]int{-2, 0, 0, 2},
		}},
		{name: "test", args: struct {
			nums   []int
			target int
		}{nums: []int{-1, 0, 1, 2, -1, -4}, target: -1}, want: [][]int{
			[]int{-4, 0, 1, 2},
			[]int{-1, -1, 0, 1},
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := fourSum(tt.args.nums, tt.args.target)

			sortOrder := func(resList [][]int) {
				// 保证结果排序唯一
				sort.Slice(resList, func(i, j int) bool {
					if resList[i][0] != resList[j][0] {
						return resList[i][0] < resList[j][0]
					}
					if resList[i][1] != resList[j][1] {
						return resList[i][1] < resList[j][1]
					}
					if resList[i][2] != resList[j][2] {
						return resList[i][2] < resList[j][2]
					}
					return resList[i][3] < resList[j][3]
				})
			}
			sortOrder(got)
			sortOrder(tt.want)

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("fourSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
