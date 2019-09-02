package leetcode

//设计一个支持 push，pop，top 操作，并能在常数时间内检索到最小元素的栈。
//
//push(x) -- 将元素 x 推入栈中。
//pop() -- 删除栈顶的元素。
//top() -- 获取栈顶元素。
//getMin() -- 检索栈中的最小元素。
//示例:
//
//MinStack minStack = new MinStack();
//minStack.push(-2);
//minStack.push(0);
//minStack.push(-3);
//minStack.getMin();   --> 返回 -3.
//minStack.pop();
//minStack.top();      --> 返回 0.
//minStack.getMin();   --> 返回 -2.

// 看了其他人的题解的开头回想起的之前看到的解法
// 因为原始数据是堆栈操作，也就导致假设当前最小值为v，若后面压入的若干个值均大于该值，则最小值不会变，同时之后一直进行pop操作，
// 直到pop到该数据前，最小值都不会变
// 同时若出现某个值大于该值，则应该更新该值为最小值，当这个值被pop后，最小值则变回原来的，也就是说最小值也是先进后出（FILO），
// 也就是说可以用堆栈来维护最小值序列，
// 新增数据时，当堆栈为空则直接压入作为最小值，否则与最小值堆栈栈顶比较，当大于该值时，无视之，小于，则压入最小值堆栈
// pop数据时，当该值与栈顶相同时，则pop最小值堆栈，否则不作操作

type MinStack struct {
	data         []int
	minElemStack []int // 辅助最小栈，当为空时一定放入，当不为空时，则仅当新的值不大于栈顶的时候才放入，当原始数据从数据栈中移除时，pop该栈
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{}
}

func (this *MinStack) Push(x int) {
	this.data = append(this.data, x)
	if len(this.minElemStack) == 0 || x <= this.minElemStack[len(this.minElemStack)-1] {
		this.minElemStack = append(this.minElemStack, x)
	}
}

func (this *MinStack) Pop() {
	top := this.Top()
	this.data = this.data[:len(this.data)-1]
	if top == this.minElemStack[len(this.minElemStack)-1] {
		this.minElemStack = this.minElemStack[:len(this.minElemStack)-1]
	}
}

func (this *MinStack) Top() int {
	return this.data[len(this.data)-1]
}

func (this *MinStack) GetMin() int {
	return this.minElemStack[len(this.minElemStack)-1]
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
