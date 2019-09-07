package leetcode

func missingNumber(nums []int) int {
	// 先利用等差数列和通项公式计算出和
	// 然后遍历一遍数组，依次减去，最后剩下的就是缺失的数
	n := len(nums)
	sum := (n + 1) * (0 + n) / 2
	for _, num := range nums {
		sum -= num
	}
	return sum
}
