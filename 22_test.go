package leetcode

import (
	"reflect"
	"testing"
)

func Test_generateParenthesis(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		struct {
			name string
			args args
			want []string
		}{name: "test", args: struct{ n int }{n: 3}, want: []string{
			"((()))",
			"(()())",
			"(())()",
			"()(())",
			"()()()",
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := generateParenthesis(tt.args.n)
			setGot := listStringToSet(got)
			setWant := listStringToSet(tt.want)
			if !reflect.DeepEqual(setGot, setWant) {
				t.Errorf("generateParenthesis() = %v, want %v", setGot, setWant)
			}
		})
	}
}
