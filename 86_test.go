package leetcode

import (
	"reflect"
	"testing"
)

// 2019/11/22 22:58 by fzls

func Test_partition(t *testing.T) {
	type args struct {
		head *ListNode
		x    int
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{name: "test", args: struct {
			head *ListNode
			x    int
		}{head: genLinkedList(1, 4, 3, 2, 5, 2), x: 3}, want: genLinkedList(1, 2, 2, 4, 3, 5)},
		{name: "test", args: struct {
			head *ListNode
			x    int
		}{head: genLinkedList(), x: 0}, want: genLinkedList()},
		{name: "test", args: struct {
			head *ListNode
			x    int
		}{head: genLinkedList(1, 4, 3, 2, 5, 2), x: 0}, want: genLinkedList(1, 4, 3, 2, 5, 2)},
		{name: "test", args: struct {
			head *ListNode
			x    int
		}{head: genLinkedList(1, 4, 3, 2, 5, 2), x: 999}, want: genLinkedList(1, 4, 3, 2, 5, 2)},
		{name: "test", args: struct {
			head *ListNode
			x    int
		}{head: genLinkedList(6, 5, 4, 3, 2, 1), x: 3}, want: genLinkedList(2, 1, 6, 5, 4, 3)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := partition(tt.args.head, tt.args.x); !reflect.DeepEqual(linkedList2List(got), linkedList2List(tt.want)) {
				t.Errorf("partition(%v) = %v, want %v", tt.args, linkedList2List(got), linkedList2List(tt.want))
			}
		})
	}
}
