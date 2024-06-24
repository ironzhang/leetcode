package leetcode

type grid struct {
	m int
	n int
}

func (g *grid) IsEnd(x, y int) bool {
	if x >= g.n-1 {
		return true
	}
	if y >= g.m-1 {
		return true
	}
	return false
}

func move(g *grid, x, y int) int {
	if g.IsEnd(x, y) {
		return 1
	}
	return move(g, x+1, y) + move(g, x, y+1)
}

func uniquePaths0(m int, n int) int {
	g := &grid{m: m, n: n}
	return move(g, 0, 0)
}

func makeGirds(m int, n int) [][]int {
	girds := make([][]int, m)
	for i := 0; i < m; i++ {
		girds[i] = make([]int, n)
	}
	return girds
}

func uniquePaths(m int, n int) int {
	girds := makeGirds(m, n)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if i == 0 || j == 0 {
				girds[i][j] = 1
				continue
			}
			girds[i][j] = girds[i-1][j] + girds[i][j-1]
		}
	}
	return girds[m-1][n-1]
}
