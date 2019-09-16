package leetcode

// 2019/09/17 0:30 by fzls
func maxArea(height []int) int {
	max := 0

	// 参考题解，使用双指针分别指向两侧
	// 若移动较大者向内，因为长度变短了，而新的高度不会高于原来的较小者，所以面积必定下降
	// 而移动较小者向内，则有可能会使面积向下，所以每次决策均移动较小者向内
	i := 0
	j := len(height) - 1
	for i < j {
		a := area(height, i, j)
		if a > max {
			max = a
		}
		if height[i] < height[j] {
			i++
		} else {
			j--
		}
	}

	return max
}

func area(heights []int, i int, j int) int {
	width := j - i
	height := heights[i]
	if height > heights[j] {
		height = heights[j]
	}
	return width * height
}
