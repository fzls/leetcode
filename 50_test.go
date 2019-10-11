package leetcode

import "testing"

func Test_myPow(t *testing.T) {
	type args struct {
		x float64
		n int
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{name: "test", args: struct {
			x float64
			n int
		}{x: 2.00000, n: 10}, want: 1024.00000},
		{name: "test", args: struct {
			x float64
			n int
		}{x: 2.10000, n: 3}, want: 9.26100},
		{name: "test", args: struct {
			x float64
			n int
		}{x: 2.00000, n: -2}, want: 0.25000},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := myPow(tt.args.x, tt.args.n); !(got-tt.want >= -1e-6 && got-tt.want <= 1e-6) {
				t.Errorf("myPow(%v) = %v, want %v", tt.args, got, tt.want)
			}
		})
	}
}
