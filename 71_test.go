package leetcode

import "testing"

// 2019/11/06 21:38 by fzls

func Test_simplifyPath(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "test", args: struct{ path string }{path: "/home/"}, want: "/home"},
		{name: "test", args: struct{ path string }{path: "/../"}, want: "/"},
		{name: "test", args: struct{ path string }{path: "/home//foo/"}, want: "/home/foo"},
		{name: "test", args: struct{ path string }{path: "/a/./b/../../c/"}, want: "/c"},
		{name: "test", args: struct{ path string }{path: "/a/../../b/../c//.//"}, want: "/c"},
		{name: "test", args: struct{ path string }{path: ""}, want: "/"},
		{name: "test", args: struct{ path string }{path: "/"}, want: "/"},
		{name: "test", args: struct{ path string }{path: "/a"}, want: "/a"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := simplifyPath(tt.args.path); got != tt.want {
				t.Errorf("simplifyPath(%v) = %v, want %v", tt.args, got, tt.want)
			}
		})
	}
}
