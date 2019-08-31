package leetcode

import "testing"

func Test_hasPathSum(t *testing.T) {
	type args struct {
		root *TreeNode
		sum  int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		struct {
			name string
			args args
			want bool
		}{name: "test", args: struct {
			root *TreeNode
			sum  int
		}{root: &TreeNode{
			5,
			&TreeNode{
				4,
				&TreeNode{
					11,
					&TreeNode{
						7,
						nil,
						nil,
					},
					&TreeNode{
						2,
						nil,
						nil,
					},
				},
				nil,
			},
			&TreeNode{
				8,
				&TreeNode{
					13,
					nil,
					nil,
				},
				&TreeNode{
					4,
					nil,
					&TreeNode{
						1,
						nil,
						nil,
					},
				},
			},
		}, sum: 22}, want: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := hasPathSum(tt.args.root, tt.args.sum); got != tt.want {
				t.Errorf("hasPathSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
