package leetcode

import (
	"reflect"
	"testing"
)

func TestFourSum(t *testing.T) {
	tests := []struct {
		nums      []int
		target    int
		solutions [][]int
	}{
		{
			nums:   []int{1, 0, -1, 0, -2, 2},
			target: 0,
			solutions: [][]int{
				{-2, -1, 1, 2},
				{-2, 0, 0, 2},
				{-1, 0, 0, 1},
			},
		},
		{
			nums:   []int{0, 0, 0, 0},
			target: 0,
			solutions: [][]int{
				{0, 0, 0, 0},
			},
		},
		{
			nums:   []int{1, 0, 0, 0},
			target: 1,
			solutions: [][]int{
				{0, 0, 0, 1},
			},
		},
		{
			nums:   []int{1, -2, -5, -4, -3, 3, 3, 5},
			target: -11,
			solutions: [][]int{
				{-5, -4, -3, 1},
			},
		},
	}
	for i, tt := range tests {
		if got, want := fourSum(tt.nums, tt.target), tt.solutions; !reflect.DeepEqual(got, want) {
			t.Errorf("%d: fourSum: got %v, want %v", i, got, want)
		}
	}
}
