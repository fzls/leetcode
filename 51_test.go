package leetcode

import (
	"reflect"
	"testing"
)

// 2019/10/24 23:05 by fzls

func Test_solveNQueens(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want [][]string
	}{
		{name: "test", args: struct{ n int }{n: 4}, want: [][]string{
			{
				".Q..", // 解法 1
				"...Q",
				"Q...",
				"..Q.",
			},
			{
				"..Q.", // 解法 2
				"Q...",
				"...Q",
				".Q..",
			},
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := solveNQueens(tt.args.n)
			sortStringPermutateList(got)
			sortStringPermutateList(tt.want)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("solveNQueens(%v) = %v, want %v", tt.args.n, got, tt.want)
			}
		})
	}
}
