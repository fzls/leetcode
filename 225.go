package leetcode

type MyStack struct {
	queue []int
	buff  []int
}

/** Initialize your data structure here. */
func Constructor_() MyStack {
	return MyStack{}
}

/** Push element x onto stack. */
func (this *MyStack) Push(x int) {
	this.queue = append(this.queue, x)
}

/** Removes the element on top of the stack and returns that element. */
func (this *MyStack) Pop() int {
	res := 0
	for len(this.queue) > 1 {
		this.buff = append(this.buff, this.queue[0])
		this.queue = this.queue[1:]
	}
	res = this.queue[0]

	this.queue = this.buff
	this.buff = nil

	return res
}

/** Get the top element. */
func (this *MyStack) Top() int {
	res := 0
	for len(this.queue) > 1 {
		this.buff = append(this.buff, this.queue[0])
		this.queue = this.queue[1:]
	}
	res = this.queue[0]
	this.buff = append(this.buff, res)

	this.queue = this.buff
	this.buff = nil

	return res
}

/** Returns whether the stack is empty. */
func (this *MyStack) Empty() bool {
	return len(this.queue) == 0
}

/**
 * Your MyStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Empty();
 */
