package leetcode

import (
	"fmt"
	"math"
)

type calculator struct {
	values map[string]int
}

func (c *calculator) sum(nums []int, i, j int) int {
	sum := 0
	key := fmt.Sprintf("%d-%d", i, j-1)
	if value, ok := c.values[key]; ok {
		sum = value + nums[j]
	} else {
		for x := i; x <= j; x++ {
			sum += nums[x]
		}
	}

	key = fmt.Sprintf("%d-%d", i, j)
	c.values[key] = sum
	return sum
}

func maxSubArrayLow(nums []int) int {
	c := calculator{values: make(map[string]int)}

	n := len(nums)
	max := math.MinInt64
	for i := 0; i < n; i++ {
		for j := i; j < n; j++ {
			s := c.sum(nums, i, j)
			if s > max {
				max = s
			}
		}
	}
	return max
}

func maxSubArrayRaw(nums []int) int {
	n := len(nums)
	sum := 0
	max := math.MinInt64
	for i := 0; i < n; i++ {
		sum += nums[i]
		if max < sum {
			max = sum
		}
		if sum < 0 {
			sum = 0
		}
	}
	return max
}

func maxSubArray(nums []int) int {
	n := len(nums)
	dp := make([]int, n)
	max := math.MinInt64
	for i := 0; i < n; i++ {
		if i == 0 {
			dp[i] = nums[i]
		} else if dp[i-1] <= 0 {
			dp[i] = nums[i]
		} else {
			dp[i] = dp[i-1] + nums[i]
		}
		if max < dp[i] {
			max = dp[i]
		}
	}
	return max
}
