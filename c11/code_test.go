package leetcode

import "testing"

func TestMaxArea(t *testing.T) {
	tests := []struct {
		height []int
		max    int
	}{
		{
			height: []int{1, 8, 6, 2, 5, 4, 8, 3, 7},
			max:    49,
		},
	}
	for i, tt := range tests {
		if got, want := maxArea(tt.height), tt.max; got != want {
			t.Errorf("%d: maxArea: got %v, want %v", i, got, want)
		}
	}
}
