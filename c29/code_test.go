package leetcode

import (
	"testing"
)

func TestDivide(t *testing.T) {
	tests := []struct {
		dividend int
		divisor  int
		division int
	}{
		{
			dividend: 10,
			divisor:  3,
			division: 3,
		},
		{
			dividend: -10,
			divisor:  -3,
			division: 3,
		},
		{
			dividend: 9,
			divisor:  3,
			division: 3,
		},
		{
			dividend: -9,
			divisor:  3,
			division: -3,
		},
		{
			dividend: 7,
			divisor:  3,
			division: 2,
		},
		{
			dividend: 7,
			divisor:  -3,
			division: -2,
		},
		{
			dividend: -2147483648,
			divisor:  -1,
			division: 2147483647,
		},
	}
	for i, tt := range tests {
		division := divide(tt.dividend, tt.divisor)
		if got, want := division, tt.division; got != want {
			t.Errorf("%d: got %v, want %v", i, got, want)
		} else {
			t.Logf("%d: got %v", i, got)
		}
	}
}
