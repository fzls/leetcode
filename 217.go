package leetcode

func containsDuplicate(nums []int) bool {
	visited := make(map[int]struct{})

	for _, num := range nums {
		if _, ok := visited[num]; ok {
			return true
		}
		visited[num] = struct{}{}
	}

	return false
}
