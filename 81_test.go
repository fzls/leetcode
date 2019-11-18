package leetcode

import "testing"

// 2019/11/18 22:38 by fzls

func Test___search(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "test", args: struct {
			nums   []int
			target int
		}{nums: []int{2, 5, 6, 0, 0, 1, 2}, target: 0}, want: true},
		{name: "test", args: struct {
			nums   []int
			target int
		}{nums: []int{2, 5, 6, 0, 0, 1, 2}, target: 3}, want: false},
		{name: "test", args: struct {
			nums   []int
			target int
		}{nums: []int{0, 0, 1, 2, 2, 5, 6}, target: 0}, want: true},
		{name: "test", args: struct {
			nums   []int
			target int
		}{nums: []int{0, 0, 1, 2, 2, 5, 6}, target: 3}, want: false},
		{name: "test", args: struct {
			nums   []int
			target int
		}{nums: []int{6, 7, 8, 1, 2, 3, 5}, target: 1}, want: true},
		{name: "test", args: struct {
			nums   []int
			target int
		}{nums: []int{6, 7, 8, 1, 2, 3, 5}, target: 4}, want: false},
		{name: "test", args: struct {
			nums   []int
			target int
		}{nums: []int{1, 2, 3, 5, 6, 7, 8}, target: 1}, want: true},
		{name: "test", args: struct {
			nums   []int
			target int
		}{nums: []int{1, 2, 3, 5, 6, 7, 8}, target: 4}, want: false},
		{name: "test", args: struct {
			nums   []int
			target int
		}{nums: []int{1}, target: 0}, want: false},
		{name: "test", args: struct {
			nums   []int
			target int
		}{nums: []int{1, 1, 3, 1}, target: 3}, want: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := __search(tt.args.nums, tt.args.target); got != tt.want {
				t.Errorf("__search(%v) = %v, want %v", tt.args, got, tt.want)
			}
		})
	}
}
