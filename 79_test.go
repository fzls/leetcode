package leetcode

import "testing"

// 2019/11/17 21:33 by fzls

func Test_exist(t *testing.T) {
	type args struct {
		board [][]byte
		word  string
	}
	board := [][]byte{
		[]byte("ABCE"),
		[]byte("SFCS"),
		[]byte("ADEE"),
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "test", args: struct {
			board [][]byte
			word  string
		}{board: board, word: "ABCCED"}, want: true},
		{name: "test", args: struct {
			board [][]byte
			word  string
		}{board: board, word: "SEE"}, want: true},
		{name: "test", args: struct {
			board [][]byte
			word  string
		}{board: board, word: "ABCB"}, want: false},
		{name: "test", args: struct {
			board [][]byte
			word  string
		}{board: board, word: ""}, want: true},
		{name: "test", args: struct {
			board [][]byte
			word  string
		}{board: board, word: "CFDE"}, want: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := exist(tt.args.board, tt.args.word); got != tt.want {
				t.Errorf("exist(%v) = %v, want %v", tt.args, got, tt.want)
			}
		})
	}
}
