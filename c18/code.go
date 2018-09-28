package leetcode

import (
	"sort"
)

func fourSum(nums []int, target int) [][]int {
	sort.Ints(nums)
	solutions := make([][]int, 0, 20)
	for i := 0; i < len(nums)-3; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		for j := i + 1; j < len(nums)-2; j++ {
			if j > i+1 && nums[j] == nums[j-1] {
				continue
			}
			x := target - (nums[i] + nums[j])
			l, r := j+1, len(nums)-1
			for l < r {
				if nums[l]+nums[r] < x {
					l = leftSideSeek(nums, l, r)
				} else if nums[l]+nums[r] > x {
					r = rightSideSeek(nums, l, r)
				} else {
					solutions = append(solutions, []int{nums[i], nums[j], nums[l], nums[r]})
					l = leftSideSeek(nums, l, r)
					r = rightSideSeek(nums, l, r)
				}
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
