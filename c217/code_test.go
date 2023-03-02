package leetcode

import "testing"

func TestContainsDuplicate(t *testing.T) {
	tests := []struct {
		nums   []int
		result bool
	}{
		{
			nums:   []int{1, 2, 3, 1},
			result: true,
		},
		{
			nums:   []int{1, 2, 3, 4},
			result: false,
		},
		{
			nums:   []int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2},
			result: true,
		},
	}
	for _, tt := range tests {
		if got, want := containsDuplicate(tt.nums), tt.result; got != want {
			t.Fatalf("%v: got %v, want %v", tt.nums, got, want)
		}
	}
}
