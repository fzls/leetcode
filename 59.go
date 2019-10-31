package leetcode

// 2019/11/01 1:09 by fzls
func generateMatrix(n int) [][]int {
	if n <= 0 {
		return nil
	}

	// init res
	buffer := make([]int, n*n)
	res := make([][]int, n)
	for i := 0; i < n; i++ {
		res[i], buffer = buffer[:n], buffer[n:]
	}

	idx := 1
	// assign value to res[row][col]; left right for col; top bottom for row;
	for step := 0; step <= (n-1)/2; step++ {
		// find left top and right bottom
		startRow, startCol := step, step
		rowSize, colSize := n-2*step, n-2*step
		endRow, endCol := startRow+rowSize-1, startCol+colSize-1

		// left[inclusive] top to right[inclusive] top
		for col := startCol; col <= endCol; col++ {
			res[startRow][col] = idx
			idx++
		}

		// right top[exclusive] to right bottom[inclusive]
		for row := startRow + 1; row <= endRow; row++ {
			res[row][endCol] = idx
			idx++
		}

		// if rowSize > 1, right[exclusive] bottom to left[inclusive] bottom
		if rowSize > 1 {
			for col := endCol - 1; col >= startCol; col-- {
				res[endRow][col] = idx
				idx++
			}
		}

		// if colSize > 1, left bottom[exclusive] to left top[exclusive]
		if colSize > 1 {
			for row := endRow - 1; row >= startRow+1; row-- {
				res[row][startCol] = idx
				idx++
			}
		}
	}

	// return
	return res
}
