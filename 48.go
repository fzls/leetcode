package leetcode

import "fmt"

// 2019/10/18 23:54 by fzls
func _rotate(matrix [][]int) {
	n := len(matrix)
	if n <= 1 {
		return
	}
	fmt.Println(matrix)

	for i := 1; i <= n/2; i++ {
		//从外到内依次处理第i圈
		for j := 0; j <= n-2; j++ {
			left := i - 1
			right := n - i
			top := i - 1
			bottom := n - i

			// 暂存右侧
			temp := matrix[top+j][right]
			// 上侧到右侧
			matrix[top+j][right] = matrix[top][left+j]
			// 左侧到上侧
			matrix[top][left+j] = matrix[bottom-j][left]
			// 下侧到左侧
			matrix[bottom-j][left] = matrix[bottom][right-j]
			// 右侧到下侧
			matrix[bottom][right-j] = temp
		}
	}

	fmt.Println(matrix)
}
