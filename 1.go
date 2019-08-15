package leetcode

func twoSum(nums []int, target int) []int {
	cache := make(map[int]int) // target - this.value => that.index

	for idx, num := range nums {
		if targetIdx, exists := cache[target-num]; exists {
			return []int{targetIdx, idx}
		}

		cache[num] = idx
	}

	return nil
}
