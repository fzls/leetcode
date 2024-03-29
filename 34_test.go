package leetcode

import (
	"reflect"
	"testing"
)

// 2019/09/29 21:51 by fzls

func Test_searchRange(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{name: "test", args: struct {
			nums   []int
			target int
		}{nums: []int{5, 7, 7, 8, 8, 10}, target: 8}, want: []int{3, 4}},
		{name: "test", args: struct {
			nums   []int
			target int
		}{nums: []int{5, 7, 7, 8, 8, 10}, target: 6}, want: []int{-1, -1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := searchRange(tt.args.nums, tt.args.target); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("searchRange() = %v, want %v", got, tt.want)
			}
		})
	}
}
