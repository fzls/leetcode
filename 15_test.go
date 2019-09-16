package leetcode

import (
	"reflect"
	"sort"
	"testing"
)

func Test_threeSum(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{name: "test", args: struct{ nums []int }{nums: []int{-1, 0, 1, 2, -1, -4}}, want: [][]int{
			[]int{-1, -1, 2},
			[]int{-1, 0, 1},
		}},
		{name: "test", args: struct{ nums []int }{nums: []int{-4, -2, -2, -2, 0, 1, 2, 2, 2, 3, 3, 4, 4, 6, 6}}, want: [][]int{
			[]int{-4, -2, 6},
			[]int{-4, 0, 4},
			[]int{-4, 1, 3},
			[]int{-4, 2, 2},
			[]int{-2, -2, 4},
			[]int{-2, 0, 2},
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := threeSum(tt.args.nums)

			sortOrder := func(resList [][]int) {
				// 保证结果排序唯一
				sort.Slice(resList, func(i, j int) bool {
					if resList[i][0] != resList[j][0] {
						return resList[i][0] < resList[j][0]
					}
					if resList[i][1] != resList[j][1] {
						return resList[i][1] < resList[j][1]
					}
					return resList[i][2] < resList[j][2]
				})
			}
			sortOrder(got)
			sortOrder(tt.want)

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("threeSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
