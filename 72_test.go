package leetcode

import "testing"

// 2019/11/07 21:39 by fzls

func Test_minDistance(t *testing.T) {
	type args struct {
		word1 string
		word2 string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "test", args: struct {
			word1 string
			word2 string
		}{word1: "horse", word2: "ros"}, want: 3},
		{name: "test", args: struct {
			word1 string
			word2 string
		}{word1: "intention", word2: "execution"}, want: 5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minDistance(tt.args.word1, tt.args.word2); got != tt.want {
				t.Errorf("minDistance(%v) = %v, want %v", tt.args, got, tt.want)
			}
		})
	}
}
