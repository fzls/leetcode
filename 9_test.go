package leetcode

import "testing"

func Test_isPalindrome(t *testing.T) {
	type args struct {
		x int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		struct {
			name string
			args args
			want bool
		}{name: "1", args: struct{ x int }{x: 121}, want: true},
		struct {
			name string
			args args
			want bool
		}{name: "2", args: struct{ x int }{x: -121}, want: false},
		struct {
			name string
			args args
			want bool
		}{name: "3", args: struct{ x int }{x: 10}, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isPalindrome(tt.args.x); got != tt.want {
				t.Errorf("isPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}