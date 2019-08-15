package leetcode

import "testing"

func Test_lengthOfLongestSubstring(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		struct {
			name string
			args args
			want int
		}{name: "case1", args: struct{ s string }{s: "abcabcbb"}, want: 3},
		struct {
			name string
			args args
			want int
		}{name: "case2", args: struct{ s string }{s: "bbbbb"}, want: 1},
		struct {
			name string
			args args
			want int
		}{name: "case3", args: struct{ s string }{s: "pwwkew"}, want: 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := lengthOfLongestSubstring(tt.args.s); got != tt.want {
				t.Errorf("lengthOfLongestSubstring() = %v, want %v", got, tt.want)
			}
		})
	}
}
