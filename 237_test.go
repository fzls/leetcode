package leetcode

import (
	"reflect"
	"testing"
)

func Test_deleteNode(t *testing.T) {
	t.Run("test1", func(t *testing.T) {
		p1 := &ListNode{
			Val: 1,
			Next: &ListNode{
				Val:  9,
				Next: nil,
			},
		}
		p5 := &ListNode{
			Val:  5,
			Next: p1,
		}
		head := &ListNode{
			Val:  4,
			Next: p5,
		}
		want := &ListNode{
			Val: 4,
			Next: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val:  9,
					Next: nil,
				},
			},
		}
		deleteNode(p5)
		if !reflect.DeepEqual(head, want) {
			t.Errorf("deleteNode() = %v, want %v", head, want)
		}
	})
	t.Run("test2", func(t *testing.T) {
		p1 := &ListNode{
			Val: 1,
			Next: &ListNode{
				Val:  9,
				Next: nil,
			},
		}
		p5 := &ListNode{
			Val:  5,
			Next: p1,
		}
		head := &ListNode{
			Val:  4,
			Next: p5,
		}
		want := &ListNode{
			Val: 4,
			Next: &ListNode{
				Val: 5,
				Next: &ListNode{
					Val:  9,
					Next: nil,
				},
			},
		}
		deleteNode(p1)
		if !reflect.DeepEqual(head, want) {
			t.Errorf("deleteNode() = %v, want %v", head, want)
		}
	})
}
