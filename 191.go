package leetcode

func hammingWeight(num uint32) int {
	// 从编程之美看到的思路
	cnt := 0
	for num != 0 {
		cnt++
		num &= (num - 1)
	}
	return cnt
}
