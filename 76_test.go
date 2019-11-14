package leetcode

import "testing"

// 2019/11/14 22:18 by fzls

func Test_minWindow(t *testing.T) {
	type args struct {
		s string
		t string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "test", args: struct {
			s string
			t string
		}{s: "ADOBECODEBANC", t: "ABC"}, want: "BANC"},
		{name: "test", args: struct {
			s string
			t string
		}{s: "ADOBECODEBANC", t: "abc"}, want: ""},
		{name: "test", args: struct {
			s string
			t string
		}{s: "ADOBECODEBANC", t: "ABCADOBECODEBANCADOBECODEBANC"}, want: ""},
		{name: "test", args: struct {
			s string
			t string
		}{s: "ADOBECODEBANC", t: "ADOBECODEBANC"}, want: "ADOBECODEBANC"},
		{name: "test", args: struct {
			s string
			t string
		}{s: "ADOBECODEBANC", t: ""}, want: ""},
		{name: "test", args: struct {
			s string
			t string
		}{s: "", t: "ABC"}, want: ""},
		{name: "test", args: struct {
			s string
			t string
		}{s: "ZADOBECODEBANCY", t: "ZY"}, want: "ZADOBECODEBANCY"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minWindow(tt.args.s, tt.args.t); got != tt.want {
				t.Errorf("minWindow(%v) = %v, want %v", tt.args, got, tt.want)
			}
		})
	}
}
