package leetcode

import "testing"

func Test_longestPalindrome(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		struct {
			name string
			args args
			want string
		}{name: "case1", args: struct{ s string }{s: "babad"}, want: "bab"},
		struct {
			name string
			args args
			want string
		}{name: "case1", args: struct{ s string }{s: "cbbd"}, want: "bb"},
		struct {
			name string
			args args
			want string
		}{name: "case1", args: struct{ s string }{s: "ab"}, want: "a"},
		struct {
			name string
			args args
			want string
		}{name: "case1", args: struct{ s string }{s: "a"}, want: "a"},
		struct {
			name string
			args args
			want string
		}{name: "case1", args: struct{ s string }{s: ""}, want: ""},
		struct {
			name string
			args args
			want string
		}{name: "case1", args: struct{ s string }{s: "bb"}, want: "bb"},
		struct {
			name string
			args args
			want string
		}{name: "case1", args: struct{ s string }{s: "ccc"}, want: "ccc"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestPalindrome(tt.args.s); got != tt.want {
				t.Errorf("longestPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}