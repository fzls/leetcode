package leetcode

func generate(numRows int) [][]int {
	var res [][]int

	for i := 0; i < numRows; i++ {
		row := make([]int, i+1)

		row[0] = 1
		row[i] = 1

		for j := 1; j < i; j++ {
			// j>=1, i>=2
			row[j] = res[i-1][j-1] + res[i-1][j]
		}

		res = append(res, row)
	}

	return res
}
