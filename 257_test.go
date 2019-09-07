package leetcode

import (
	"reflect"
	"testing"
)

func Test_binaryTreePaths(t *testing.T) {
	type args struct {
		root *TreeNode
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
		}{name: "test", args: struct{ root *TreeNode }{root: &TreeNode{
			Val: 1,
			Left: &TreeNode{
				Val:  2,
				Left: nil,
				Right: &TreeNode{
					Val:   5,
					Left:  nil,
					Right: nil,
				},
			},
			Right: &TreeNode{
				Val:   3,
				Left:  nil,
				Right: nil,
			},
		}}, want: []string{"1->2->5", "1->3"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := binaryTreePaths(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("binaryTreePaths() = %v, want %v", got, tt.want)
			}
		})
	}
}
