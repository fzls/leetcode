package leetcode

func isPowerOfTwo(n int) bool {
	//等价于判断n的二进制表示中仅有1个1
	cntOfOne := 0
	for n != 0 {
		cntOfOne++
		n = n & (n - 1)

		if cntOfOne > 1 {
			return false
		}
	}
	return cntOfOne == 1
}
