package leetcode

import "testing"

func Test__rotate(t *testing.T) {
	type args struct {
		matrix [][]int
	}
	tests := []struct {
		name string
		args args
	}{
		struct {
			name string
			args args
		}{name: "test", args: struct{ matrix [][]int }{matrix: [][]int{
			{1, 2, 3},
			{4, 5, 6},
			{7, 8, 9},
		}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_rotate(tt.args.matrix)
		})
	}
}
