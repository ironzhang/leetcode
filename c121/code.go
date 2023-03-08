package leetcode

import "math"

func minPrices(prices []int) []int {
	min := math.MaxInt64
	values := make([]int, len(prices))
	for i, p := range prices {
		if p < min {
			min = p
		}
		values[i] = min
	}
	return values
}

func maxProfit(prices []int) int {
	max := 0
	mins := minPrices(prices)
	for i, p := range prices {
		profit := p - mins[i]
		if profit > max {
			max = profit
		}
	}
	return max
}
