package leetcode

import "testing"

func Test_isPalindrome__(t *testing.T) {
	type args struct {
		head *ListNode
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "test", args: struct{ head *ListNode }{head: &ListNode{
			Val: 1,
			Next: &ListNode{
				Val:  2,
				Next: nil,
			},
		}}, want: false}, {name: "test", args: struct{ head *ListNode }{head: &ListNode{
			Val: 1,
			Next: &ListNode{
				Val: 2,
				Next: &ListNode{
					Val: 2,
					Next: &ListNode{
						Val:  1,
						Next: nil,
					},
				},
			},
		}}, want: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isPalindrome__(tt.args.head); got != tt.want {
				t.Errorf("isPalindrome__() = %v, want %v", got, tt.want)
			}
		})
	}
}
