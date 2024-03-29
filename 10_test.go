package leetcode

import "testing"

func Test_isMatch(t *testing.T) {
	type args struct {
		s string
		p string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "test", args: struct {
			s string
			p string
		}{s: "aa", p: "a"}, want: false},
		{name: "test", args: struct {
			s string
			p string
		}{s: "aa", p: "a*"}, want: true},
		{name: "test", args: struct {
			s string
			p string
		}{s: "ab", p: ".*"}, want: true},
		{name: "test", args: struct {
			s string
			p string
		}{s: "aab", p: "c*a*b"}, want: true},
		{name: "test", args: struct {
			s string
			p string
		}{s: "mississippi", p: "mis*is*p*."}, want: false},
		{name: "test", args: struct {
			s string
			p string
		}{s: "a", p: "ab"}, want: false},
		{name: "test", args: struct {
			s string
			p string
		}{s: "ab", p: ".*c"}, want: false},
		{name: "test", args: struct {
			s string
			p string
		}{s: "aaa", p: "ab*ac*a"}, want: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isMatch(tt.args.s, tt.args.p); got != tt.want {
				t.Errorf("isMatch() = %v, want %v", got, tt.want)
			}
		})
	}
}
