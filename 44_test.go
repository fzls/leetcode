package leetcode

import "testing"

func Test__isMatch(t *testing.T) {
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
		}{s: "aa", p: "*"}, want: true},
		{name: "test", args: struct {
			s string
			p string
		}{s: "cb", p: "?a"}, want: false},
		{name: "test", args: struct {
			s string
			p string
		}{s: "adceb", p: "*a*b"}, want: true},
		{name: "test", args: struct {
			s string
			p string
		}{s: "acdcb", p: "a*c?b"}, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := __isMatch(tt.args.s, tt.args.p); got != tt.want {
				t.Errorf("_isMatch(%v) = %v, want %v", tt.args, got, tt.want)
			}
		})
	}
}
