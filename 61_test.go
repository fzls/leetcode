package leetcode

import (
	"reflect"
	"testing"
)

// 2019/11/02 23:51 by fzls

func Test_rotateRight(t *testing.T) {
	type args struct {
		head *ListNode
		k    int
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{name: "test", args: struct {
			head *ListNode
			k    int
		}{head: genLinkedList(1, 2, 3, 4, 5), k: 2}, want: genLinkedList(4, 5, 1, 2, 3)},
		{name: "test", args: struct {
			head *ListNode
			k    int
		}{head: genLinkedList(0, 1, 2), k: 4}, want: genLinkedList(2, 0, 1)},
		{name: "test", args: struct {
			head *ListNode
			k    int
		}{head: genLinkedList(0, 1, 2), k: 3}, want: genLinkedList(0, 1, 2)},
		{name: "test", args: struct {
			head *ListNode
			k    int
		}{head: genLinkedList(0), k: 3}, want: genLinkedList(0)},
		{name: "test", args: struct {
			head *ListNode
			k    int
		}{head: genLinkedList(1, 2), k: 3}, want: genLinkedList(2, 1)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := rotateRight(tt.args.head, tt.args.k); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("rotateRight() = %v, want %v", got, tt.want)
			}
		})
	}
}
