package leetcode

func maxProfit_(prices []int) int {
	if len(prices) == 0 {
		return 0
	}

	// 转化价格数组为价格变动数组
	for i := len(prices) - 1; i > 0; i-- {
		prices[i] -= prices[i-1]
	}
	prices[0] = 0

	mp := 0

	for i := 0; i < len(prices); i++ {
		if prices[i] > 0 {
			mp += prices[i]
		}
	}

	return mp
}
