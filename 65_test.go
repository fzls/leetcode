package leetcode

import "testing"

// 2019/11/03 22:07 by fzls

func Test_isNumber(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "test", args: struct{ s string }{s: "0"}, want: true},
		{name: "test", args: struct{ s string }{s: "01"}, want: true},
		{name: "test", args: struct{ s string }{s: "01.1"}, want: true},
		{name: "test", args: struct{ s string }{s: "10.1"}, want: true},
		{name: "test", args: struct{ s string }{s: "10.01"}, want: true},
		{name: "test", args: struct{ s string }{s: "0.1"}, want: true},
		{name: "test", args: struct{ s string }{s: ".1"}, want: true},
		{name: "test", args: struct{ s string }{s: "1."}, want: true},
		{name: "test", args: struct{ s string }{s: "."}, want: false},
		{name: "test", args: struct{ s string }{s: "abc"}, want: false},
		{name: "test", args: struct{ s string }{s: "1 a"}, want: false},
		{name: "test", args: struct{ s string }{s: "2e10"}, want: true},
		{name: "test", args: struct{ s string }{s: " -90e3"}, want: true},
		{name: "test", args: struct{ s string }{s: " 1e"}, want: false},
		{name: "test", args: struct{ s string }{s: "e3"}, want: false},
		{name: "test", args: struct{ s string }{s: " 6e-1"}, want: true},
		{name: "test", args: struct{ s string }{s: " 99e2.5"}, want: false},
		{name: "test", args: struct{ s string }{s: "53.5e93"}, want: true},
		{name: "test", args: struct{ s string }{s: " --6"}, want: false},
		{name: "test", args: struct{ s string }{s: "-+3"}, want: false},
		{name: "test", args: struct{ s string }{s: "95a54e53"}, want: false},

		{name: "test", args: struct{ s string }{s: ""}, want: false},
		{name: "test", args: struct{ s string }{s: "   -12.345e-678    "}, want: true},
		{name: "test", args: struct{ s string }{s: "   -12.e-678    "}, want: true},
		{name: "test", args: struct{ s string }{s: "   -.345e-678    "}, want: true},
		{name: "test", args: struct{ s string }{s: "-12.345e-678    "}, want: true},
		{name: "test", args: struct{ s string }{s: "   12.345e-678    "}, want: true},
		{name: "test", args: struct{ s string }{s: "   -12e-678    "}, want: true},
		{name: "test", args: struct{ s string }{s: "   -12.345e678    "}, want: true},
		{name: "test", args: struct{ s string }{s: "   -12.345    "}, want: true},
		{name: "test", args: struct{ s string }{s: "   -12.345e-678"}, want: true},
		{name: "test", args: struct{ s string }{s: "   -12.e-678    "}, want: true},

		{name: "test", args: struct{ s string }{s: "   -12.345e-    "}, want: false},
		{name: "test", args: struct{ s string }{s: "   -12.345e    "}, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isNumber(tt.args.s); got != tt.want {
				t.Errorf("isNumber(%v) = %v, want %v", tt.args.s, got, tt.want)
			}
		})
	}
}
