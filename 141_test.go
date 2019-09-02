package leetcode

import "testing"

func Test_hasCycle(t *testing.T) {
	type args struct {
		head *ListNode
	}

	ring1Tail := &ListNode{
		Val:  -4,
		Next: nil,
	}
	ring1 := &ListNode{
		Val: 2,
		Next: &ListNode{
			Val:  0,
			Next: ring1Tail,
		},
	}
	ring1Tail.Next = ring1
	list1 := &ListNode{
		Val:  3,
		Next: ring1,
	}

	ring2Tail := &ListNode{
		Val:  2,
		Next: nil,
	}
	ring2 := &ListNode{
		Val:  1,
		Next: ring2Tail,
	}
	ring2Tail.Next = ring2
	list2 := ring2

	list3 := &ListNode{
		Val:  1,
		Next: nil,
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "test", args: struct{ head *ListNode }{head: list1}, want: true},
		{name: "test", args: struct{ head *ListNode }{head: list2}, want: true},
		{name: "test", args: struct{ head *ListNode }{head: list3}, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := hasCycle(tt.args.head); got != tt.want {
				t.Errorf("hasCycle() = %v, want %v", got, tt.want)
			}
		})
	}
}
