package leetcode

import "testing"

func Test_convert(t *testing.T) {
	type args struct {
		s       string
		numRows int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		struct {
			name string
			args args
			want string
		}{name: "case1", args: struct {
			s       string
			numRows int
		}{s: "LEETCODEISHIRING", numRows: 3}, want: "LCIRETOESIIGEDHN"},
		struct {
			name string
			args args
			want string
		}{name: "case2", args: struct {
			s       string
			numRows int
		}{s: "LEETCODEISHIRING", numRows: 4}, want: "LDREOEIIECIHNTSG"},
		struct {
			name string
			args args
			want string
		}{name: "case3", args: struct {
			s       string
			numRows int
		}{s: "A", numRows: 1}, want: "A"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := convert(tt.args.s, tt.args.numRows); got != tt.want {
				t.Errorf("convert() = %v, want %v", got, tt.want)
			}
		})
	}
}