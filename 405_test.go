package leetcode

import "testing"

func Test_toHex(t *testing.T) {
	type args struct {
		num int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "test", args: struct{ num int }{num: 26}, want: "1a"},
		{name: "test", args: struct{ num int }{num: -1}, want: "ffffffff"},
		{name: "test", args: struct{ num int }{num: 0}, want: "0"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := toHex(tt.args.num); got != tt.want {
				t.Errorf("toHex() = %v, want %v", got, tt.want)
			}
		})
	}
}
