package leetcode

import (
	"reflect"
	"testing"
)

func Test_removeNthFromEnd(t *testing.T) {
	type args struct {
		head *ListNode
		n    int
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{name: "test", args: struct {
			head *ListNode
			n    int
		}{head: genLinkedList(1, 2, 3, 4, 5), n: 2}, want: genLinkedList(1, 2, 3, 5)},
		{name: "test", args: struct {
			head *ListNode
			n    int
		}{head: genLinkedList(1, 2), n: 1}, want: genLinkedList(1)},
		{name: "test", args: struct {
			head *ListNode
			n    int
		}{head: genLinkedList(1), n: 1}, want: genLinkedList()},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := removeNthFromEnd(tt.args.head, tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("removeNthFromEnd() = %v, want %v", linkedList2List(got), linkedList2List(tt.want))
			}
		})
	}
}
