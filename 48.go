package leetcode

// 2019/10/18 23:54 by fzls
func _rotate(matrix [][]int) {
	n := len(matrix)
	if n <= 1 {
		return
	}

	// 从外到内遍历每一圈
	for level := 1; level <= n/2; level++ {
		rowMin := level - 1
		rowMax := n - level
		colMin := level - 1
		colMax := n - level
		// 将每一圈分为四块相等的区域，每个区域的元素数目为 n-2*(level-1)-1
		batchSize := n - 2*(level-1) - 1
		for offset := 0; offset < batchSize; offset++ {
			// 在缓存区中放入上区块的元素
			temp := matrix[rowMin][colMin+offset]
			// 在上区块放入左区块的元素
			matrix[rowMin][colMin+offset] = matrix[rowMax-offset][colMin]
			// 在左区块放入下区块的元素
			matrix[rowMax-offset][colMin] = matrix[rowMax][colMax-offset]
			// 在下区块放入右区块的元素
			matrix[rowMax][colMax-offset] = matrix[rowMin+offset][colMax]
			// 在右区块放入缓存区的元素
			matrix[rowMin+offset][colMax] = temp
		}
	}
}
