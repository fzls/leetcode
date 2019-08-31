package leetcode

import (
	"reflect"
	"testing"
)

func Test_plusOne(t *testing.T) {
	type args struct {
		digits []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{name: "test", args: struct{ digits []int }{digits: []int{1, 2, 3}}, want: []int{1, 2, 4}},
		{name: "test", args: struct{ digits []int }{digits: []int{4, 3, 2, 1}}, want: []int{4, 3, 2, 2}},
		{name: "test", args: struct{ digits []int }{digits: []int{9, 9, 9}}, want: []int{1, 0, 0, 0}},
		{name: "test", args: struct{ digits []int }{digits: []int{0}}, want: []int{1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := plusOne(tt.args.digits); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("plusOne() = %v, want %v", got, tt.want)
			}
		})
	}
}
