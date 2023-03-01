package leetcode

func largestLocal(grid [][]int) [][]int {
	n := len(grid)
	maxLocal := make([][]int, n-2)
	for i := 0; i < n-2; i++ {
		maxLocal[i] = make([]int, n-2)
		for j := 0; j < n-2; j++ {
			maxLocal[i][j] = maxGrid(grid, i+1, j+1)
		}
	}
	return maxLocal
}

func maxGrid(grid [][]int, x, y int) int {
	max := 0
	for i := x - 1; i <= x+1; i++ {
		for j := y - 1; j <= y+1; j++ {
			if grid[i][j] > max {
				max = grid[i][j]
			}
		}
	}
	return max
}
