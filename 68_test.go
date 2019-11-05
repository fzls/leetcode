package leetcode

import (
	"reflect"
	"testing"
)

// 2019/11/06 1:18 by fzls

func Test_fullJustify(t *testing.T) {
	type args struct {
		words    []string
		maxWidth int
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{name: "test", args: struct {
			words    []string
			maxWidth int
		}{words: []string{"This", "is", "an", "example", "of", "text", "justification."}, maxWidth: 16}, want: []string{
			"This    is    an",
			"example  of text",
			"justification.  ",
		}},
		{name: "test", args: struct {
			words    []string
			maxWidth int
		}{words: []string{"What", "must", "be", "acknowledgment", "shall", "be"}, maxWidth: 16}, want: []string{
			"What   must   be",
			"acknowledgment  ",
			"shall be        ",
		}},
		{name: "test", args: struct {
			words    []string
			maxWidth int
		}{words: []string{"a"}, maxWidth: 16}, want: []string{
			"a               ",
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := fullJustify(tt.args.words, tt.args.maxWidth); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("fullJustify(%v) \n   = %v, \nwant %v", tt.args, got, tt.want)
			}
		})
	}
}
