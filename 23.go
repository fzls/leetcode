package leetcode

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
func (lh *ListHeap) Push(x *ListNode) {
	// 放到末尾并上移
	lh.data = append(lh.data, x)
	lh.up(len(lh.data) - 1)
}

func (lh *ListHeap) Pop() *ListNode {
	// 将头部取出作为结果，并将原先的最后一个放到该处，然后下移
	top := lh.data[0]

	lh.data[0] = lh.data[len(lh.data)-1]
	lh.data = lh.data[:len(lh.data)-1]
	lh.down(0)

	return top
}

func (lh *ListHeap) down(i int) {
	n := len(lh.data)
	for i < n {
		// 找到左右子节点中的较小者
		// 左节点
		left := 2*i + 1

		if left >= n {
			break
		}
		minChild := left
		// 右节点
		if right := 2*i + 2; right < n && lh.data[right].Val < lh.data[left].Val {
			minChild = right
		}
		// 判断是否比子节点小，若是则无需移动
		if lh.data[i].Val <= lh.data[minChild].Val {
			break
		}

		// 否则交换位置，继续
		lh.data[i], lh.data[minChild] = lh.data[minChild], lh.data[i]
		i = minChild
	}
}

func (lh *ListHeap) up(i int) {
	for i != 0 {
		// 找到父节点
		parent := (i - 1) / 2

		// 判断是否需要上移
		if lh.data[i].Val < lh.data[parent].Val {
			lh.data[i], lh.data[parent] = lh.data[parent], lh.data[i]
			i = parent
		} else {
			break
		}
	}
}

func mergeKLists(lists []*ListNode) *ListNode {
	dummy := &ListNode{}

	// heap的写法记不太清了，先用官方提供的最小堆实现吧- -
	listHeap := &ListHeap{}

	for _, head := range lists {
		if head != nil {
			listHeap.Push(head)
		}
	}

	current := dummy
	for listHeap.Len() != 0 {
		minHead := listHeap.Pop()
		current.Next = minHead
		current = current.Next

		if minHead.Next != nil {
			listHeap.Push(minHead.Next)
		}
	}

	return dummy.Next
}
