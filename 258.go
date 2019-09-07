package leetcode

func addDigits(num int) int {
	// 参考题解，可以发现每次变化数目都是9的倍数，最终结果是1-9中的一个
	// 先把原数减一，这样取余得到0-8，然后加回去，就得到了1-9
	return (num-1)%9 + 1
}
