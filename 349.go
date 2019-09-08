package leetcode

func intersection(nums1 []int, nums2 []int) []int {
	set1 := make(map[int]struct{})
	for _, num := range nums1 {
		set1[num] = struct{}{}
	}

	resSet := make(map[int]struct{})
	for _, num := range nums2 {
		if _, ok := set1[num]; ok {
			resSet[num] = struct{}{}
		}
	}

	var res []int
	for num := range resSet {
		res = append(res, num)
	}
	return res
}
