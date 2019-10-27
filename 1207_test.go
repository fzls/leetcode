package leetcode

import "testing"

// 2019/10/27 20:40 by fzls

func Test_uniqueOccurrences(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "test", args: struct{ arr []int }{arr: []int{1, 2, 2, 1, 1, 3}}, want: true},
		{name: "test", args: struct{ arr []int }{arr: []int{1, 2}}, want: false},
		{name: "test", args: struct{ arr []int }{arr: []int{-3, 0, 1, -3, 1, 1, 1, -3, 10, 0}}, want: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := uniqueOccurrences(tt.args.arr); got != tt.want {
				t.Errorf("uniqueOccurrences(%v) = %v, want %v", tt.args, got, tt.want)
			}
		})
	}
}
