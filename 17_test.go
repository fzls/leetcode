package leetcode

import (
	"reflect"
	"testing"
)

func Test_letterCombinations(t *testing.T) {
	type args struct {
		digits string
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
		}{name: "test", args: struct{ digits string }{digits: "23"}, want: []string{"ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := letterCombinations(tt.args.digits)
			gotSet := listStringToSet(got)
			wantSet := listStringToSet(tt.want)
			if !reflect.DeepEqual(gotSet, wantSet) {
				t.Errorf("letterCombinations() = %v, want %v", gotSet, wantSet)
			}
		})
	}
}
