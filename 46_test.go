package leetcode

import (
	"reflect"
	"testing"
)

func Test_permute(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{name: "test", args: struct{ nums []int }{nums: []int{1, 2, 3}}, want: [][]int{
			{1, 2, 3},
			{1, 3, 2},
			{2, 1, 3},
			{2, 3, 1},
			{3, 1, 2},
			{3, 2, 1},
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := permute(tt.args.nums)
			sortIntPermutateList(got)
			sortIntPermutateList(tt.want)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("permute(%v) = %v, want %v", tt.args.nums, got, tt.want)
			}
		})
	}
}
