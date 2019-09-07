package leetcode

import "testing"

func Test_wordPattern(t *testing.T) {
	type args struct {
		pattern string
		str     string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "test", args: struct {
			pattern string
			str     string
		}{pattern: "abba", str: "dog cat cat dog"}, want: true},
		{name: "test", args: struct {
			pattern string
			str     string
		}{pattern: "abba", str: "dog cat cat fish"}, want: false},
		{name: "test", args: struct {
			pattern string
			str     string
		}{pattern: "aaaa", str: "dog cat cat dog"}, want: false},
		{name: "test", args: struct {
			pattern string
			str     string
		}{pattern: "abba", str: "dog dog dog dog"}, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := wordPattern(tt.args.pattern, tt.args.str); got != tt.want {
				t.Errorf("wordPattern() = %v, want %v", got, tt.want)
			}
		})
	}
}
