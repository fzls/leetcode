package leetcode

// 2019/10/24 22:42 by fzls
func solveNQueens(n int) [][]string {
	qb := NewQueenBoard(n)
	qb.Solve()

	return qb.Results
}

type QueenBoard struct {
	N     int
	Board [][]bool

	Rows        []bool
	Colomns     []bool
	LeftAngles  []bool // /
	RightAngles []bool // \

	Results [][]string
}

func NewQueenBoard(n int) *QueenBoard {
	board := make([][]bool, n)
	data := make([]bool, n*n)
	for i := 0; i < n; i++ {
		board[i], data = data[:n], data[n:]
	}
	return &QueenBoard{
		N:           n,
		Board:       board,
		Rows:        make([]bool, n),
		Colomns:     make([]bool, n),
		LeftAngles:  make([]bool, 2*n-1),
		RightAngles: make([]bool, 2*n-1),
	}
}

func (qb *QueenBoard) CanPut(row, col int) bool {
	return !(qb.Rows[row] || qb.Colomns[col] || qb.LeftAngles[row+col] || qb.RightAngles[row-col+(qb.N-1)])
}

func (qb *QueenBoard) Put(row, col int) {
	qb.Board[row][col] = true
	qb.Rows[row] = true
	qb.Colomns[col] = true
	qb.LeftAngles[row+col] = true
	qb.RightAngles[row-col+(qb.N-1)] = true
}

func (qb *QueenBoard) Remove(row, col int) {
	qb.Board[row][col] = false
	qb.Rows[row] = false
	qb.Colomns[col] = false
	qb.LeftAngles[row+col] = false
	qb.RightAngles[row-col+(qb.N-1)] = false
}

func (qb *QueenBoard) Solve() {
	qb.solveCore(0)
}

func (qb *QueenBoard) solveCore(row int) {
	if row == qb.N {
		// 得到一个合法的结果，保存
		var res []string
		for i := 0; i < qb.N; i++ {
			line := make([]byte, qb.N)
			for j := 0; j < qb.N; j++ {
				if qb.Board[i][j] == true {
					line[j] = 'Q'
				} else {
					line[j] = '.'
				}
			}
			res = append(res, string(line))
		}
		qb.Results = append(qb.Results, res)
		return
	}

	// TODO：看看能否剪枝

	for col := 0; col < qb.N; col++ {
		if !qb.CanPut(row, col) {
			continue
		}

		qb.Put(row, col)
		qb.solveCore(row + 1)
		qb.Remove(row, col)
	}
}
