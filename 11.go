package leetcode

// 2019/09/17 0:30 by fzls
func maxArea(height []int) int {
	max := 0

	// 暴力解法
	for i := 0; i < len(height); i++ {
		for j := i + 1; j < len(height); j++ {
			a := area(height, i, j)
			if a > max {
				max = a
			}
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
