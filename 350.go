package leetcode

func intersect(nums1 []int, nums2 []int) []int {
	cntMap := make(map[int]int)
	for _, num := range nums1 {
		cntMap[num]++
	}

	var res []int
	for _, num := range nums2 {
		if cntMap[num] > 0 {
			res = append(res, num)
			cntMap[num]--
		}
	}

	return res
}
