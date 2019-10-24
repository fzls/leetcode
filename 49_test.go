package leetcode

import (
	"reflect"
	"testing"
)

func Test_groupAnagrams(t *testing.T) {
	type args struct {
		strs []string
	}
	tests := []struct {
		name string
		args args
		want [][]string
	}{
		struct {
			name string
			args args
			want [][]string
		}{name: "test", args: struct{ strs []string }{strs: []string{
			"eat", "tea", "tan", "ate", "nat", "bat",
		}}, want: [][]string{
			{"ate", "eat", "tea"},
			{"nat", "tan"},
			{"bat"},
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := groupAnagrams(tt.args.strs)
			sortStringList(got)
			sortStringList(tt.want)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("groupAnagrams() \nres  %v, \nwant %v", got, tt.want)
			}
		})
	}
}
