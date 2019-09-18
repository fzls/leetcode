package leetcode

type _Stack struct {
	data []int
}

func (this *_Stack) Push(x int) {
	this.data = append(this.data, x)
}

func (this *_Stack) Pop() int {
	top := this.Top()
	this.data = this.data[:len(this.data)-1]
	return top
}

func (this *_Stack) Top() int {
	return this.data[len(this.data)-1]
}

func (this *_Stack) Empty() bool {
	return len(this.data) == 0
}

type MyQueue struct {
	stack _Stack
	buff  _Stack
}

/** Initialize your data structure here. */
func Constructor__() MyQueue {
	return MyQueue{}
}

/** Push element x to the back of queue. */
func (this *MyQueue) Push(x int) {
	this.stack.Push(x)
}

/** Removes the element from in front of queue and returns that element. */
func (this *MyQueue) Pop() int {
	for !this.stack.Empty() {
		this.buff.Push(this.stack.Pop())
	}
	queueFront := this.buff.Pop()
	for !this.buff.Empty() {
		this.stack.Push(this.buff.Pop())
	}
	return queueFront
}

/** Get the front element. */
func (this *MyQueue) Peek() int {
	for !this.stack.Empty() {
		this.buff.Push(this.stack.Pop())
	}
	queueFront := this.buff.Top()
	for !this.buff.Empty() {
		this.stack.Push(this.buff.Pop())
	}
	return queueFront
}

/** Returns whether the queue is empty. */
func (this *MyQueue) Empty() bool {
	return this.stack.Empty()
}

/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */
