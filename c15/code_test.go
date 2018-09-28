package leetcode

import (
	"reflect"
	"testing"
)

func TestThreeSum(t *testing.T) {
	tests := []struct {
		nums []int
		set  [][]int
	}{
		{
			nums: []int{-1, 0, 1, 2, -1, -4},
			set: [][]int{
				{-1, -1, 2},
				{-1, 0, 1},
			},
		},
		{
			nums: []int{0, 0, 0},
			set: [][]int{
				{0, 0, 0},
			},
		},
		{
			nums: []int{1, 0, 0, 0},
			set: [][]int{
				{0, 0, 0},
			},
		},
		{
			nums: []int{3, 0, -2, -1, 1, 2},
			set: [][]int{
				{-2, -1, 3},
				{-2, 0, 2},
				{-1, 0, 1},
			},
		},
	}
	for i, tt := range tests {
		if got, want := threeSum(tt.nums), tt.set; !reflect.DeepEqual(got, want) {
			t.Errorf("%d: threeSum: got %v, want %v", i, got, want)
		}
	}
}
