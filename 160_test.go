package leetcode

import (
	"reflect"
	"testing"
)

func Test_getIntersectionNode(t *testing.T) {
	type args struct {
		headA *ListNode
		headB *ListNode
	}

	common1 := &ListNode{
		Val: 8,
		Next: &ListNode{
			Val: 4,
			Next: &ListNode{
				Val:  5,
				Next: nil,
			},
		},
	}
	ha1 := &ListNode{
		Val: 4,
		Next: &ListNode{
			Val:  1,
			Next: common1,
		},
	}
	hb1 := &ListNode{
		Val: 5,
		Next: &ListNode{
			Val: 0,
			Next: &ListNode{
				Val:  1,
				Next: common1,
			},
		},
	}

	common2 := &ListNode{
		Val: 2,
		Next: &ListNode{
			Val:  4,
			Next: nil,
		},
	}
	ha2 := &ListNode{
		Val: 0,
		Next: &ListNode{
			Val: 9,
			Next: &ListNode{
				Val:  1,
				Next: common2,
			},
		},
	}
	hb2 := &ListNode{
		Val:  3,
		Next: common2,
	}

	var common3 *ListNode = nil
	ha3 := &ListNode{
		Val: 2,
		Next: &ListNode{
			Val: 6,
			Next: &ListNode{
				Val:  4,
				Next: common3,
			},
		},
	}
	hb3 := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val:  5,
			Next: common3,
		},
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{name: "test", args: struct {
			headA *ListNode
			headB *ListNode
		}{headA: ha1, headB: hb1}, want: common1},
		{name: "test", args: struct {
			headA *ListNode
			headB *ListNode
		}{headA: ha2, headB: hb2}, want: common2},
		{name: "test", args: struct {
			headA *ListNode
			headB *ListNode
		}{headA: ha3, headB: hb3}, want: common3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getIntersectionNode(tt.args.headA, tt.args.headB); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getIntersectionNode() = %v, want %v", got, tt.want)
			}
		})
	}
}
