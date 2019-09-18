package leetcode

func getSum(a int, b int) int {
	// 自己没想出来，看他人题解得到的思路
	for a != 0 && b != 0 {
		a, b = a^b, a&b<<1
	}

	if a != 0 {
		return a
	} else {
		return b
	}
}
