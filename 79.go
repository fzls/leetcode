package leetcode

// 2019/11/17 21:09 by fzls
func exist(board [][]byte, word string) bool {
	if len(word) == 0 {
		return true
	}
	m := len(board)
	if m == 0 {
		return false
	}
	n := len(board[0])
	if n == 0 {
		return false
	}

	visited := make([][]bool, m)
	for i := 0; i < m; i++ {
		visited[i] = make([]bool, n)
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if dfsExist(board, i, j, word, 0, visited) {
				return true
			}
		}
	}

	return false
}

type Direction struct {
	Row int
	Col int
}

var dirs = [...]Direction{
	{-1, 0}, // up
	{1, 0},  // down
	{0, -1}, // left
	{0, 1},  // right
}

func dfsExist(board [][]byte, row, col int, word string, idx int, visited [][]bool) bool {
	m, n := len(board), len(board[0])

	if board[row][col] != word[idx] {
		return false
	}

	if idx == len(word)-1 {
		return true
	}

	visited[row][col] = true

	for _, dir := range dirs {
		nextRow, nextCol := row+dir.Row, col+dir.Col
		if nextRow >= 0 && nextRow <= m-1 && nextCol >= 0 && nextCol <= n-1 && !visited[nextRow][nextCol] {
			if dfsExist(board, nextRow, nextCol, word, idx+1, visited) {
				return true
			}
		}
	}

	visited[row][col] = false

	return false
}
