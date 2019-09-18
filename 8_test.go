package leetcode

import "testing"

func Test_myAtoi(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		struct {
			name string
			args args
			want int
		}{name: "1", args: struct{ str string }{str: "42"}, want: 42},
		struct {
			name string
			args args
			want int
		}{name: "2", args: struct{ str string }{str: "   -42"}, want: -42},
		struct {
			name string
			args args
			want int
		}{name: "3", args: struct{ str string }{str: "4193 with words"}, want: 4193},
		struct {
			name string
			args args
			want int
		}{name: "4", args: struct{ str string }{str: "words and 987"}, want: 0},
		struct {
			name string
			args args
			want int
		}{name: "5", args: struct{ str string }{str: "-91283472332"}, want: -2147483648},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := myAtoi(tt.args.str); got != tt.want {
				t.Errorf("myAtoi() = %v, want %v", got, tt.want)
			}
		})
	}
}