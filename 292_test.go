package leetcode

import "testing"

func Test_canWinNim(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "test", args: struct{ n int }{n: 4}, want: false},
		{name: "test", args: struct{ n int }{n: 1}, want: true},
		{name: "test", args: struct{ n int }{n: 2}, want: true},
		{name: "test", args: struct{ n int }{n: 3}, want: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := canWinNim(tt.args.n); got != tt.want {
				t.Errorf("canWinNim() = %v, want %v", got, tt.want)
			}
		})
	}
}
