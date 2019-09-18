package leetcode

import "testing"

func Test_longestCommonPrefix(t *testing.T) {
	type args struct {
		strs []string
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
		}{name: "1", args: struct{ strs []string }{strs: []string{"flower", "flow", "flight"}}, want: "fl"},
		struct {
			name string
			args args
			want string
		}{name: "2", args: struct{ strs []string }{strs: []string{"dog","racecar","car"}}, want: ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestCommonPrefix(tt.args.strs); got != tt.want {
				t.Errorf("longestCommonPrefix() = %v, want %v", got, tt.want)
			}
		})
	}
}
