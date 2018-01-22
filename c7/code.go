package leetcode

import "math"

func abs(x int) (sign int, y int) {
	if x < 0 {
		return -1, -x
	}
	return 1, x
}

func reverse(x int) int {
	if x < math.MinInt32 || x > math.MaxInt32 {
		return 0
	}

	sum := 0
	sign, y := abs(x)
	for y > 0 {
		sum = sum*10 + y%10
		y /= 10
	}
	x = sign * sum

	if x < math.MinInt32 || x > math.MaxInt32 {
		return 0
	}
	return x
}
