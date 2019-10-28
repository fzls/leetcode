package leetcode

import (
	"reflect"
	"testing"
)

// 2019/10/29 0:38 by fzls

func Test_spiralOrder(t *testing.T) {
	type args struct {
		matrix [][]int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{name: "test", args: struct{ matrix [][]int }{matrix: [][]int{
			{1, 2, 3},
			{4, 5, 6},
			{7, 8, 9},
		}}, want: []int{
			1, 2, 3, 6, 9, 8, 7, 4, 5,
		}},
		{name: "test", args: struct{ matrix [][]int }{matrix: [][]int{
			{1, 2, 3, 4},
			{5, 6, 7, 8},
			{9, 10, 11, 12},
		}}, want: []int{
			1, 2, 3, 4, 8, 12, 11, 10, 9, 5, 6, 7,
		}},
		{name: "test", args: struct{ matrix [][]int }{matrix: [][]int{}}, want: nil},
		{name: "test", args: struct{ matrix [][]int }{matrix: [][]int{
			{6, 9, 7},
		}}, want: []int{
			6, 9, 7,
		}},
		{name: "test", args: struct{ matrix [][]int }{matrix: [][]int{
			{6},
			{9},
			{7},
		}}, want: []int{
			6, 9, 7,
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := spiralOrder(tt.args.matrix); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("spiralOrder(%v) = %v, want %v", tt.args, got, tt.want)
			}
		})
	}
}
