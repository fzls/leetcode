package leetcode

import "testing"

func Test_trap(t *testing.T) {
	type args struct {
		height []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "test", args: struct{ height []int }{height: []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}}, want: 6},
		{name: "test", args: struct{ height []int }{height: []int{5, 4, 1, 2}}, want: 1},
		{name: "test", args: struct{ height []int }{height: []int{5, 2, 1, 2, 1, 5}}, want: 14},
		{name: "test", args: struct{ height []int }{height: []int{2, 8, 5, 5, 6, 1, 7, 4, 5}}, want: 12},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := trap(tt.args.height); got != tt.want {
				t.Errorf("trap(%v) = %v, want %v", tt.args.height, got, tt.want)
			}
		})
	}
}
