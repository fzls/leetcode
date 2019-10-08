package leetcode

import (
	"reflect"
	"testing"
)

func Test_permuteUnique(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{name: "test", args: struct{ nums []int }{nums: []int{1, 1, 2}}, want: [][]int{
			{1, 1, 2},
			{1, 2, 1},
			{2, 1, 1},
		}},
		{name: "test", args: struct{ nums []int }{nums: []int{0, 1, 0, 0, 9}}, want: [][]int{
			{0, 0, 0, 1, 9},
			{0, 0, 0, 9, 1},
			{0, 0, 1, 0, 9},
			{0, 0, 1, 9, 0},
			{0, 0, 9, 0, 1},
			{0, 0, 9, 1, 0},
			{0, 1, 0, 0, 9},
			{0, 1, 0, 9, 0},
			{0, 1, 9, 0, 0},
			{0, 9, 0, 0, 1},
			{0, 9, 0, 1, 0},
			{0, 9, 1, 0, 0},
			{1, 0, 0, 0, 9},
			{1, 0, 0, 9, 0},
			{1, 0, 9, 0, 0},
			{1, 9, 0, 0, 0},
			{9, 0, 0, 0, 1},
			{9, 0, 0, 1, 0},
			{9, 0, 1, 0, 0},
			{9, 1, 0, 0, 0},
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := permuteUnique(tt.args.nums)
			sortIntPermutateList(got)
			sortIntPermutateList(tt.want)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("permuteUnique(%v) =\nres  %v, \nwant %v", tt.args.nums, got, tt.want)
			}
		})
	}
}
