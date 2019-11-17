package leetcode

import (
	"reflect"
	"testing"
)

// 2019/11/17 21:00 by fzls

func Test_subsets(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{name: "test", args: struct{ nums []int }{nums: []int{1, 2, 3}}, want: [][]int{
			{3},
			{1},
			{2},
			{1, 2, 3},
			{1, 3},
			{2, 3},
			{1, 2},
			{},
		}},
		{name: "test", args: struct{ nums []int }{nums: []int{1, 2}}, want: [][]int{
			{},
			{1},
			{2},
			{1, 2},
		}},
		{name: "test", args: struct{ nums []int }{nums: []int{1}}, want: [][]int{
			{},
			{1},
		}},
		{name: "test", args: struct{ nums []int }{nums: []int{}}, want: [][]int{
			{},
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := subsets(tt.args.nums)
			sortIntListList(got)
			sortIntListList(tt.want)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("subsets(%v) \nres= %v, \nwant %v", tt.args.nums, got, tt.want)
			}
		})
	}
}
