package leetcode

import "testing"

func TestRomanToInt(t *testing.T) {
	tests := []struct {
		i int
		r string
	}{
		{
			i: 3,
			r: "III",
		},
		{
			i: 4,
			r: "IV",
		},
		{
			i: 9,
			r: "IX",
		},
		{
			i: 58,
			r: "LVIII",
		},
		{
			i: 1994,
			r: "MCMXCIV",
		},
		{
			i: 1094,
			r: "MXCIV",
		},
		{
			i: 621,
			r: "DCXXI",
		},
	}
	for i, tt := range tests {
		if got, want := romanToInt(tt.r), tt.i; got != want {
			t.Errorf("%d: romanToInt: got %v, want %v", i, got, want)
		}
	}
}
