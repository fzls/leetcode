package leetcode

import "testing"

func Test_mySqrt(t *testing.T) {
	type args struct {
		x int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "test", args: struct{ x int }{x: 4}, want: 2},
		{name: "test", args: struct{ x int }{x: 8}, want: 2},
		{name: "test", args: struct{ x int }{x: 1}, want: 1},
		{name: "test", args: struct{ x int }{x: 9}, want: 3},
		{name: "test", args: struct{ x int }{x: 10}, want: 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mySqrt(tt.args.x); got != tt.want {
				t.Errorf("mySqrt() = %v, want %v", got, tt.want)
			}
		})
	}
}
