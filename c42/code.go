package leetcode

func minInt(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func maxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func lookupLeftSideMaxHeight(height []int, n int) int {
	maxHeight := 0
	for i := n - 1; i >= 0; i-- {
		if height[i] > maxHeight {
			maxHeight = height[i]
		}
	}
	return maxHeight
}

func lookupRightSideMaxHeight(height []int, n int) int {
	maxHeight := 0
	for i := n + 1; i < len(height); i++ {
		if height[i] > maxHeight {
			maxHeight = height[i]
		}
	}
	return maxHeight
}

func trap0(height []int) int {
	sum := 0
	for i, h := range height {
		leftMax := lookupLeftSideMaxHeight(height, i)
		rightMax := lookupRightSideMaxHeight(height, i)
		y := minInt(leftMax, rightMax)
		n := y - h
		if n > 0 {
			sum += n
		}
	}
	return sum
}

func makeMaxLeftSideHeights(height []int) []int {
	n := len(height)
	maxs := make([]int, n)
	for i := 0; i < n; i++ {
		if i == 0 {
			maxs[i] = 0
		} else {
			maxs[i] = maxInt(maxs[i-1], height[i-1])
		}
	}
	return maxs
}

func makeMaxRightSideHeights(height []int) []int {
	n := len(height)
	maxs := make([]int, n)
	for i := n - 1; i >= 0; i-- {
		if i == n-1 {
			maxs[i] = 0
		} else {
			maxs[i] = maxInt(maxs[i+1], height[i+1])
		}
	}
	return maxs
}

func trap(height []int) int {
	maxLeftSideHeights := makeMaxLeftSideHeights(height)
	maxRightSideHeights := makeMaxRightSideHeights(height)

	sum := 0
	for i, h := range height {
		y := minInt(maxLeftSideHeights[i], maxRightSideHeights[i])
		n := y - h
		if n > 0 {
			sum += n
		}
	}
	return sum
}
