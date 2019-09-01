package leetcode

import "testing"

func Test_isPalindrome_(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "test", args: struct{ s string }{s: "A man, a plan, a canal: Panama"}, want: true},
		{name: "test", args: struct{ s string }{s: "race a car"}, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isPalindrome_(tt.args.s); got != tt.want {
				t.Errorf("isPalindrome_() = %v, want %v", got, tt.want)
			}
		})
	}
}
