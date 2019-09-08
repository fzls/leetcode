package leetcode

import "testing"

func Test_getSum(t *testing.T) {
	type args struct {
		a int
		b int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "test", args: struct {
			a int
			b int
		}{a: 1, b: 2}, want: 3},
		{name: "test", args: struct {
			a int
			b int
		}{a: -2, b: 3}, want: 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getSum(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("getSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
