package leetcode

import "testing"

// 2019/09/28 1:18 by fzls

func Test_longestValidParentheses(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "test", args: struct{ s string }{s: "(()"}, want: 2},
		{name: "test", args: struct{ s string }{s: ")()())"}, want: 4},
		{name: "test", args: struct{ s string }{s: "(()()("}, want: 4},
		{name: "test", args: struct{ s string }{s: "(()())"}, want: 6},
		{name: "test", args: struct{ s string }{s: "()()()((((((())"}, want: 6},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestValidParentheses(tt.args.s); got != tt.want {
				t.Errorf("longestValidParentheses(%v) = %v, want %v", tt.args.s, got, tt.want)
			}
		})
	}
}
