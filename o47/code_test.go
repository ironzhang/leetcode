package leetcode

import "testing"

func TestMaxValue(t *testing.T) {
	tests := []struct {
		grid [][]int
		max  int
	}{
		{
			grid: [][]int{
				{1, 3, 1},
				{1, 5, 1},
				{4, 2, 1},
			},
			max: 12,
		},
	}
	for _, tt := range tests {
		if got, want := maxValue(tt.grid), tt.max; got != want {
			t.Fatalf("max values: got %v, want %v", got, want)
		}
	}
}
