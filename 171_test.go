package leetcode

import "testing"

func Test_titleToNumber(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "test", args: struct{ s string }{s: "A"}, want: 1},
		{name: "test", args: struct{ s string }{s: "AB"}, want: 28},
		{name: "test", args: struct{ s string }{s: "ZY"}, want: 701},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := titleToNumber(tt.args.s); got != tt.want {
				t.Errorf("titleToNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}
