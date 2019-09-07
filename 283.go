package leetcode

func _swap(nums []int, i, j int) {
	temp := nums[i]
	nums[i] = nums[j]
	nums[j] = temp
}

func moveZeroes(nums []int) {
	// 使用冒泡法，将所有不是0的都往前移
	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			continue
		}

		for j := i - 1; j >= 0; j-- {
			if nums[j] != 0 {
				break
			}
			_swap(nums, j, j+1)
		}
	}

}
