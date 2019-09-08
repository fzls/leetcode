package leetcode

import "testing"

func Test_isPowerOfThree(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "test", args: struct{ n int }{n: 27}, want: true},
		{name: "test", args: struct{ n int }{n: 0}, want: false},
		{name: "test", args: struct{ n int }{n: 9}, want: true},
		{name: "test", args: struct{ n int }{n: 45}, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isPowerOfThree(tt.args.n); got != tt.want {
				t.Errorf("isPowerOfThree() = %v, want %v", got, tt.want)
			}
		})
	}
}
