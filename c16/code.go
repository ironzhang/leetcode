package leetcode

import (
	"sort"
)

func threeSumClosest(nums []int, target int) int {
	if len(nums) < 3 {
		panic("invalid params")
	}

	sort.Ints(nums)
	closest := nums[0] + nums[1] + nums[2]
	for i := 0; i < len(nums)-2; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		l, r := i+1, len(nums)-1
		for l < r {
			x := nums[i] + nums[l] + nums[r]
			if abs(x-target) < abs(closest-target) {
				closest = x
			}
			if x < target {
				l = leftSideSeek(nums, l, r)
			} else if x > target {
				r = rightSideSeek(nums, l, r)
			} else {
				break
			}
		}
	}
	return closest
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func leftSideSeek(nums []int, l, r int) int {
	for l++; l < r && nums[l] == nums[l-1]; l++ {
	}
	return l
}

func rightSideSeek(nums []int, l, r int) int {
	for r--; l < r && nums[r] == nums[r+1]; r-- {
	}
	return r
}
