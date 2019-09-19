package leetcode

import (
	"reflect"
	"testing"
)

func Test_mergeKLists(t *testing.T) {
	type args struct {
		lists []*ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		struct {
			name string
			args args
			want *ListNode
		}{name: "test", args: struct{ lists []*ListNode }{lists: []*ListNode{
			genLinkedList(1, 4, 5),
			genLinkedList(1, 3, 4),
			genLinkedList(2, 6),
		}}, want: genLinkedList(1, 1, 2, 3, 4, 4, 5, 6)},
		struct {
			name string
			args args
			want *ListNode
		}{name: "test", args: struct{ lists []*ListNode }{lists: []*ListNode{
			genLinkedList(),
		}}, want: genLinkedList()},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mergeKLists(tt.args.lists); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("mergeKLists() = %v, want %v", linkedList2List(got), linkedList2List(tt.want))
			}
		})
	}
}
