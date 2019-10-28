package leetcode

// 2019/10/29 0:16 by fzls
func spiralOrder(matrix [][]int) []int {
	if len(matrix) == 0 {
		return nil
	}

	if len(matrix[0]) == 0 {
		return nil
	}

	m, n := len(matrix), len(matrix[0])
	res := make([]int, 0, m*n)

	startX, startY := 0, 0

	step := 0
	for {
		endX, endY := startX+m-2*step-1, startY+n-2*step-1
		if endX < startX || endY < startY {
			break
		}

		// [x, y] -> [x, y + n-2*step -1]
		for y := startY; y <= endY; y++ {
			res = append(res, matrix[startX][y])
		}

		// [x+1, y+ n-2*step -1] -> [x + m-2*step-1, y+ n-2*step - 1]
		for x := startX + 1; x <= endX; x++ {
			res = append(res, matrix[x][endY])
		}

		// [x + m-2*step -1, y + n-2*step-1 -1] -> [x + m-2*step - 1, y]
		if startX != endX {
			for y := endY - 1; y >= startY; y-- {
				res = append(res, matrix[endX][y])
			}
		}

		// [x + m-2*step-1 -1, y] -> [x + 1, y]
		if startY != endY {
			for x := endX - 1; x >= startX+1; x-- {
				res = append(res, matrix[x][startY])
			}
		}

		step++
		startX++
		startY++
	}

	return res
}
