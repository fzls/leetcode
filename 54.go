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
		if endX == startX {
			for y := startY; y <= endY; y++ {
				res = append(res, matrix[startX][y])
			}
			break
		}
		if endY == startY {
			for x := startX; x <= endX; x++ {
				res = append(res, matrix[x][startY])
			}
			break
		}

		for y := startY; y <= endY-1; y++ {
			res = append(res, matrix[startX][y])
		}

		for x := startX; x <= endX-1; x++ {
			res = append(res, matrix[x][endY])
		}

		for y := endY; y >= startY+1; y-- {
			res = append(res, matrix[endX][y])
		}

		for x := endX; x >= startX+1; x-- {
			res = append(res, matrix[x][startY])
		}

		step++
		startX++
		startY++
	}

	return res
}
