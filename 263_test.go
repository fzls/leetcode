package leetcode

import "testing"

func Test_isUgly(t *testing.T) {
	type args struct {
		num int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "test", args: struct{ num int }{num: 6}, want: true},
		{name: "test", args: struct{ num int }{num: 8}, want: true},
		{name: "test", args: struct{ num int }{num: 14}, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isUgly(tt.args.num); got != tt.want {
				t.Errorf("isUgly() = %v, want %v", got, tt.want)
			}
		})
	}
}
