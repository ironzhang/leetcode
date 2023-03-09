package leetcode

func matrixReshape(mat [][]int, r int, c int) [][]int {
	m := len(mat)
	n := len(mat[0])

	if m*n != r*c {
		return mat
	}

	index := 0
	newmat := makeMatrix(r, c)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			x := index / c
			y := index % c
			newmat[x][y] = mat[i][j]
			index++
		}
	}
	return newmat
}

func makeMatrix(r, c int) [][]int {
	m := make([][]int, r)
	for i := 0; i < r; i++ {
		m[i] = make([]int, c)
	}
	return m
}
