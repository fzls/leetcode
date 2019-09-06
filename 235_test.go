package leetcode

import (
	"reflect"
	"testing"
)

func Test_lowestCommonAncestor(t *testing.T) {
	type args struct {
		root *TreeNode
		p    *TreeNode
		q    *TreeNode
	}
	p1 := &TreeNode{
		Val: 2,
		Left: &TreeNode{
			Val:   0,
			Left:  nil,
			Right: nil,
		},
		Right: &TreeNode{
			Val: 4,
			Left: &TreeNode{
				Val:   3,
				Left:  nil,
				Right: nil,
			},
			Right: &TreeNode{
				Val:   5,
				Left:  nil,
				Right: nil,
			},
		},
	}
	q1 := &TreeNode{
		Val: 8,
		Left: &TreeNode{
			Val:   7,
			Left:  nil,
			Right: nil,
		},
		Right: &TreeNode{
			Val:   9,
			Left:  nil,
			Right: nil,
		},
	}
	c1 := &TreeNode{
		Val:   6,
		Left:  p1,
		Right: q1,
	}
	root1 := c1

	q2 := &TreeNode{
		Val: 4,
		Left: &TreeNode{
			Val:   3,
			Left:  nil,
			Right: nil,
		},
		Right: &TreeNode{
			Val:   5,
			Left:  nil,
			Right: nil,
		},
	}
	p2 := &TreeNode{
		Val: 2,
		Left: &TreeNode{
			Val:   0,
			Left:  nil,
			Right: nil,
		},
		Right: q2,
	}
	c2 := p2
	root2 := &TreeNode{
		Val:  6,
		Left: c2,
		Right: &TreeNode{
			Val: 8,
			Left: &TreeNode{
				Val:   7,
				Left:  nil,
				Right: nil,
			},
			Right: &TreeNode{
				Val:   9,
				Left:  nil,
				Right: nil,
			},
		},
	}

	p3 := &TreeNode{
		Val:   3,
		Left:  nil,
		Right: nil,
	}
	q3 := &TreeNode{
		Val:   5,
		Left:  nil,
		Right: nil,
	}
	c3 := &TreeNode{
		Val:   4,
		Left:  p3,
		Right: q3,
	}
	root3 := &TreeNode{
		Val: 6,
		Left: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val:   0,
				Left:  nil,
				Right: nil,
			},
			Right: c3,
		},
		Right: &TreeNode{
			Val: 8,
			Left: &TreeNode{
				Val:   7,
				Left:  nil,
				Right: nil,
			},
			Right: &TreeNode{
				Val:   9,
				Left:  nil,
				Right: nil,
			},
		},
	}

	tests := []struct {
		name string
		args args
		want *TreeNode
	}{
		{name: "test", args: struct {
			root *TreeNode
			p    *TreeNode
			q    *TreeNode
		}{root: root1, p: p1, q: q1}, want: c1},
		{name: "test", args: struct {
			root *TreeNode
			p    *TreeNode
			q    *TreeNode
		}{root: root2, p: p2, q: q2}, want: c2},
		{name: "test", args: struct {
			root *TreeNode
			p    *TreeNode
			q    *TreeNode
		}{root: root3, p: p3, q: q3}, want: c3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := lowestCommonAncestor(tt.args.root, tt.args.p, tt.args.q); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("lowestCommonAncestor() = %v, want %v", got, tt.want)
			}
		})
	}
}
