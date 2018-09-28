package leetcode

import (
	"sort"
)

func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	solutions := make([][]int, 0, 20)
	for i := 0; i < len(nums)-2; i++ {
		if nums[i] > 0 {
			break
		}
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		x := -nums[i]
		l, r := i+1, len(nums)-1
		for l < r {
			if nums[l]+nums[r] < x {
				l = leftSideSeek(nums, l, r)
			} else if nums[l]+nums[r] > x {
				r = rightSideSeek(nums, l, r)
			} else {
				solutions = append(solutions, []int{nums[i], nums[l], nums[r]})
				l = leftSideSeek(nums, l, r)
				r = rightSideSeek(nums, l, r)
			}
		}
	}
	return solutions
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
