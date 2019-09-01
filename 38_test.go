package leetcode

import "testing"

func Test_countAndSay(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "test", args: struct{ n int }{n: 1}, want: "1"},
		{name: "test", args: struct{ n int }{n: 2}, want: "11"},
		{name: "test", args: struct{ n int }{n: 3}, want: "21"},
		{name: "test", args: struct{ n int }{n: 4}, want: "1211"},
		{name: "test", args: struct{ n int }{n: 5}, want: "111221"},
		{name: "test", args: struct{ n int }{n: 6}, want: "312211"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countAndSay(tt.args.n); got != tt.want {
				t.Errorf("countAndSay() = %v, want %v", got, tt.want)
			}
		})
	}
}