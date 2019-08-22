package leetcode

import "math"

func abs(x int) (sign int, y int) {
	if x < 0 {
		return -1, -x
	}
	return 1, x
}

func max(x int) (maxsum, maxpop int) {
	if x < 0 {
		return (math.MaxInt32 + 1) / 10, (math.MaxInt32 + 1) % 10
	}
	return (math.MaxInt32) / 10, (math.MaxInt32) % 10
}

func reverse(x int) int {
	if x < math.MinInt32 || x > math.MaxInt32 {
		return 0
	}
	sum := 0
	sign, y := abs(x)
	maxsum, maxpop := max(x)
	for y > 0 {
		pop := y % 10
		y /= 10
		if sum > maxsum || (sum == maxsum && pop > maxpop) {
			return 0
		}
		sum = sum*10 + pop
	}
	return sign * sum
}
