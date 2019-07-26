package leetcode

const boardSize = 9
const subBoxSize = 3

func isValidSudoku(board [][]byte) bool {
	for i := 0; i < boardSize; i++ {
		if !isValidRow(board, i) {
			return false
		}
		if !isValidColumn(board, i) {
			return false
		}
	}
	for i := 0; i < boardSize; i += subBoxSize {
		for j := 0; j < boardSize; j += subBoxSize {
			if !isValidSubBox(board, i, j) {
				return false
			}
		}
	}
	return true
}

func isValidRow(board [][]byte, r int) bool {
	var m [10]int
	for i := 0; i < boardSize; i++ {
		if ch := board[r][i]; ch != '.' {
			k := ch - '0'
			if m[k] > 0 {
				return false
			}
			m[k]++
		}
	}
	return true
}

func isValidColumn(board [][]byte, c int) bool {
	var m [10]int
	for i := 0; i < boardSize; i++ {
		if ch := board[i][c]; ch != '.' {
			k := ch - '0'
			if m[k] > 0 {
				return false
			}
			m[k]++
		}
	}
	return true
}

func isValidSubBox(board [][]byte, r, c int) bool {
	var m [10]int
	for i := r; i < r+subBoxSize; i++ {
		for j := c; j < c+subBoxSize; j++ {
			if ch := board[i][j]; ch != '.' {
				k := ch - '0'
				if m[k] > 0 {
					return false
				}
				m[k]++
			}
		}
	}
	return true
}
