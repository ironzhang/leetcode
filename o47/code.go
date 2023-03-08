package leetcode

// maxValues[i,j] = grid[i,j] + max(maxValues[i-1,j], maxValues[i,j-1])
func maxValue(grid [][]int) int {
	m := len(grid)
	n := len(grid[0])

	maxValues := make([][]int, m)
	for i := 0; i < m; i++ {
		maxValues[i] = make([]int, n)
		for j := 0; j < n; j++ {
			max := getValue(maxValues, i-1, j)
			if value := getValue(maxValues, i, j-1); value > max {
				max = value
			}
			maxValues[i][j] = grid[i][j] + max
		}
	}

	return maxValues[m-1][n-1]
}

func getValue(values [][]int, i, j int) int {
	if i < 0 || j < 0 {
		return 0
	}
	return values[i][j]
}
