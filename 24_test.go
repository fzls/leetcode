package leetcode

import (
	"reflect"
	"testing"
)

func Test_swapPairs(t *testing.T) {
	type args struct {
		head *ListNode
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
		}{name: "test", args: struct{ head *ListNode }{head: genLinkedList(1, 2, 3, 4)}, want: genLinkedList(2, 1, 4, 3)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := swapPairs(tt.args.head); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("swapPairs() = %v, want %v", linkedList2List(got), linkedList2List(tt.want))
			}
		})
	}
}
