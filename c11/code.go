package leetcode

func maxArea(height []int) int {
	max := 0
	for i := 0; i < len(height)-1; i++ {
		ai := height[i]
		for j := i + 1; j < len(height); j++ {
			area := calcArea(i, ai, j, height[j])
			if area > max {
				max = area
			}
		}
	}
	return max
}

func calcArea(i, ai, j, aj int) int {
	if ai < aj {
		return (j - i) * ai
	} else {
		return (j - i) * aj
	}
}
