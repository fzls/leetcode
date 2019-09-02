package leetcode

func singleNumber(nums []int) int {
	cntMap := make(map[int]int)
	for _, num := range nums {
		if cntMap[num] == 1 {
			delete(cntMap, num)
		} else {
			cntMap[num] = 1
		}
	}
	for k, _ := range cntMap {
		return k
	}

	return 0
}
