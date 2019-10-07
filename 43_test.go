package leetcode

import "testing"

// 2019/10/08 0:52 by fzls

func Test_multiply(t *testing.T) {
	type args struct {
		num1 string
		num2 string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "test", args: struct {
			num1 string
			num2 string
		}{num1: "2", num2: "3"}, want: "6"},
		{name: "test", args: struct {
			num1 string
			num2 string
		}{num1: "123", num2: "456"}, want: "56088"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := multiply(tt.args.num1, tt.args.num2); got != tt.want {
				t.Errorf("multiply(%v) = %v, want %v", tt.args, got, tt.want)
			}
		})
	}
}
