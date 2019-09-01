package leetcode

import "testing"

func Test_romanToInt(t *testing.T) {
	type args struct {
		s string
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
		}{name: "1", args: struct{ s string }{s: "III"}, want: 3},
		struct {
			name string
			args args
			want int
		}{name: "2", args: struct{ s string }{s: "IV"}, want: 4},
		struct {
			name string
			args args
			want int
		}{name: "3", args: struct{ s string }{s: "IX"}, want: 9},
		struct {
			name string
			args args
			want int
		}{name: "4", args: struct{ s string }{s: "LVIII"}, want: 58},
		struct {
			name string
			args args
			want int
		}{name: "5", args: struct{ s string }{s: "MCMXCIV"}, want: 1994},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := romanToInt(tt.args.s); got != tt.want {
				t.Errorf("romanToInt() = %v, want %v", got, tt.want)
			}
		})
	}
}