package leetcode

import "testing"

func TestMaxProfit(t *testing.T) {
	tests := []struct {
		prices []int
		max    int
	}{
		{
			prices: []int{7, 1, 5, 3, 6, 4},
			max:    5,
		},
		{
			prices: []int{7, 6, 4, 3, 1},
			max:    0,
		},
	}
	for _, tt := range tests {
		if got, want := maxProfit(tt.prices), tt.max; got != want {
			t.Fatalf("%v's max profit: got %v, want %v", tt.prices, got, want)
		}
	}
}
