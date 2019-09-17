package leetcode

// 2019/09/18 0:23 by fzls
func listIntToSet(nums []int) map[int]struct{} {
	set := make(map[int]struct{}, len(nums))
	for _, num := range nums {
		set[num] = struct{}{}
	}
	return set
}

func listStringToSet(nums []string) map[string]struct{} {
	set := make(map[string]struct{}, len(nums))
	for _, num := range nums {
		set[num] = struct{}{}
	}
	return set
}
