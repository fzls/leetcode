package leetcode

import "container/heap"

// 2019/09/20 0:47 by fzls
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
type ListHeap struct {
	data []*ListNode
}

func (lh *ListHeap) Len() int {
	return len(lh.data)
}

func (lh *ListHeap) Less(i, j int) bool {
	return lh.data[i].Val < lh.data[j].Val
}

func (lh *ListHeap) Swap(i, j int) {
	lh.data[i], lh.data[j] = lh.data[j], lh.data[i]
}

func (lh *ListHeap) Push(x interface{}) {
	lh.data = append(lh.data, x.(*ListNode))
}

func (lh *ListHeap) Pop() interface{} {
	top := lh.data[len(lh.data)-1]
	lh.data = lh.data[:len(lh.data)-1]
	return top
}

func mergeKLists(lists []*ListNode) *ListNode {
	dummy := &ListNode{}

	// heap的写法记不太清了，先用官方提供的最小堆实现吧- -
	listHeap := &ListHeap{}
	heap.Init(listHeap)

	for _, head := range lists {
		if head != nil {
			heap.Push(listHeap, head)
		}
	}

	current := dummy
	for listHeap.Len() != 0 {
		minHead := heap.Pop(listHeap).(*ListNode)
		current.Next = minHead
		current = current.Next

		if minHead.Next != nil {
			heap.Push(listHeap, minHead.Next)
		}
	}

	return dummy.Next
}
