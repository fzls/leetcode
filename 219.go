package leetcode

func containsNearbyDuplicate(nums []int, k int) bool {
	lastIndexMap := make(map[int]int)
	for idx, num := range nums {
		if lastIdx, ok := lastIndexMap[num]; ok && idx-lastIdx <= k {
			return true
		}
		lastIndexMap[num] = idx
	}

	return false
}
