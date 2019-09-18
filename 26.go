package leetcode

//func Run(){
//	removeDuplicates([]int{0,0,1,1,1,2,2,3,3,4})
//}

func removeDuplicates(nums []int) int {
	if len(nums) <= 1 {
		return len(nums)
	}

	i := 0
	j := 1
	n := len(nums)
	for i < n && j < n {
		if nums[j] != nums[i] {
			i++
			nums[i] = nums[j]
		}
		j++
	}

	return i + 1
}
