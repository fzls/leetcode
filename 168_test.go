package leetcode

import "testing"

func Test_convertToTitle(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "test", args: struct{ n int }{n: 1}, want: "A"},
		{name: "test", args: struct{ n int }{n: 28}, want: "AB"},
		{name: "test", args: struct{ n int }{n: 701}, want: "ZY"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := convertToTitle(tt.args.n); got != tt.want {
				t.Errorf("convertToTitle() = %v, want %v", got, tt.want)
			}
		})
	}
}
