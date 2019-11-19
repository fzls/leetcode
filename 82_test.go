package leetcode

import (
	"reflect"
	"testing"
)

// 2019/11/19 22:18 by fzls

func Test__deleteDuplicates(t *testing.T) {
	type args struct {
		head *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{name: "test", args: struct{ head *ListNode }{head: genLinkedList(1, 2, 3, 3, 4, 4, 5)}, want: genLinkedList(1, 2, 5)},
		{name: "test", args: struct{ head *ListNode }{head: genLinkedList(1, 1, 1, 2, 3)}, want: genLinkedList(2, 3)},
		{name: "test", args: struct{ head *ListNode }{head: genLinkedList(1, 1, 1)}, want: genLinkedList()},
		{name: "test", args: struct{ head *ListNode }{head: genLinkedList(1, 2, 2)}, want: genLinkedList(1)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := _deleteDuplicates(tt.args.head); !reflect.DeepEqual(linkedList2List(got), linkedList2List(tt.want)) {
				t.Errorf("_deleteDuplicates(%v) = %v, want %v", linkedList2List(tt.args.head), linkedList2List(got), linkedList2List(tt.want))
			}
		})
	}
}
