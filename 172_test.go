package leetcode

import "testing"

func Test_trailingZeroes(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "test", args: struct{ n int }{n: 3}, want: 0},
		{name: "test", args: struct{ n int }{n: 5}, want: 1},
		{name: "test", args: struct{ n int }{n: 0}, want: 0},
		{name: "test", args: struct{ n int }{n: 6}, want: 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := trailingZeroes(tt.args.n); got != tt.want {
				t.Errorf("trailingZeroes() = %v, want %v", got, tt.want)
			}
		})
	}
}
