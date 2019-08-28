package leetcode

import "testing"

func Test_isValid(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "1", args: struct{ s string }{s: "()"}, want: true},
		{name: "2", args: struct{ s string }{s:  "()[]{}"}, want: true},
		{name: "3", args: struct{ s string }{s: "(]"}, want: false},
		{name: "4", args: struct{ s string }{s: "([)]"}, want: false},
		{name: "5", args: struct{ s string }{s: "{[]}"}, want: true},
		{name: "5", args: struct{ s string }{s: "]"}, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isValid(tt.args.s); got != tt.want {
				t.Errorf("isValid() = %v, want %v", got, tt.want)
			}
		})
	}
}