package leetcode

import "testing"

func Test_addStrings(t *testing.T) {
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
		}{num1: "3456", num2: "128"}, want: "3584"},
		{name: "test", args: struct {
			num1 string
			num2 string
		}{num1: "3456", num2: "0"}, want: "3456"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := addStrings(tt.args.num1, tt.args.num2); got != tt.want {
				t.Errorf("addStrings() = %v, want %v", got, tt.want)
			}
		})
	}
}
