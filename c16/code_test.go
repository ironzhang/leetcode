package leetcode

import (
	"testing"
)

func TestThreeSum(t *testing.T) {
	tests := []struct {
		nums    []int
		target  int
		closest int
	}{
		{
			nums:    []int{-1, 2, 1, -4},
			target:  1,
			closest: 2,
		},
		{
			nums:    []int{-1, 0, 1, 2, -1, -4},
			target:  -4,
			closest: -4,
		},
		{
			nums:    []int{0, 0, 0},
			target:  10,
			closest: 0,
		},
		{
			nums:    []int{1, 0, 0, 0},
			target:  -1,
			closest: 0,
		},
		{
			nums:    []int{1, 0, 0, 0},
			target:  5,
			closest: 1,
		},
		{
			nums:    []int{3, 0, -2, -1, 1, 2},
			target:  -1,
			closest: -1,
		},
	}
	for i, tt := range tests {
		if got, want := threeSumClosest(tt.nums, tt.target), tt.closest; got != want {
			t.Errorf("%d: threeSumClosest: got %v, want %v", i, got, want)
		}
	}
}
