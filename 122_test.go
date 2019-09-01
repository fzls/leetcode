package leetcode

import "testing"

func Test_maxProfit_(t *testing.T) {
	type args struct {
		prices []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "test", args: struct{ prices []int }{prices: []int{7, 1, 5, 3, 6, 4}}, want: 7},
		{name: "test", args: struct{ prices []int }{prices: []int{1, 2, 3, 4, 5}}, want: 4},
		{name: "test", args: struct{ prices []int }{prices: []int{7, 6, 4, 3, 1}}, want: 0},
		{name: "test", args: struct{ prices []int }{prices: []int{}}, want: 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxProfit_(tt.args.prices); got != tt.want {
				t.Errorf("maxProfit_() = %v, want %v", got, tt.want)
			}
		})
	}
}
