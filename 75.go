package leetcode

// 2019/11/10 23:57 by fzls
func sortColors(nums []int) {
	var cnts [3]int
	for _, num := range nums {
		cnts[num]++
	}

	index := 0
	for color := 0; color <= 2; color++ {
		for cnt := 0; cnt < cnts[color]; cnt++ {
			nums[index] = color
			index++
		}
	}

	return
}
