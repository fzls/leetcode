package leetcode

func isPowerOfTwo(n int) bool {
	//等价于判断n的二进制表示中仅有1个1
	return n > 0 && n&(n-1) == 0
}
