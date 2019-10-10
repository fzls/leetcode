package leetcode

// 2019/10/10 23:40 by fzls
func trap(height []int) int {
	rain := 0

	left := 0
	for left < len(height)-1 {
		// 从当前为止找到第一个大于右侧一格的地方作为本次的池子的左侧
		for left < len(height)-1 && !(height[left] > height[left+1]) {
			left++
		}
		if left >= len(height)-1 {
			break
		}

		right := left + 1
		// 继续向右搜索找到第一个大于左侧一格的位置作为本次的池子的右侧开始
		for right < len(height) && !(height[right] > height[right-1]) {
			right++
		}
		if right == len(height) {
			break
		}
		// 如果右侧大于本位置，则继续向右扩展右侧边缘
		for right < len(height)-1 && height[right+1] >= height[right] {
			right++
		}
		// 如果继续往右扩展能找到一个大于等于左侧位置的地方，则直接将其更新位右侧边缘
		for idx := right; idx < len(height); idx++ {
			if idx != right && height[idx] >= height[right] {
				right = idx
			}
			if height[idx] >= height[left] {
				right = idx
				break
			}
		}
		// 计算从left到right的水面高度，然后计算中间池子大小
		rainHeight := height[left]
		if rainHeight > height[right] {
			rainHeight = height[right]
		}
		//fmt.Printf("from %v to %v, height = %v\n", left, right, rainHeight)
		for i := left + 1; i < right; i++ {
			if rainHeight-height[i] > 0 {
				rain += rainHeight - height[i]
				//fmt.Println(i, "=>", rainHeight-height[i])
			}
		}

		left = right
	}
	//fmt.Println("-------------------------------------------")

	return rain
}
