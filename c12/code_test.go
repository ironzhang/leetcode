package leetcode

import "testing"

func TestIntToRoman(t *testing.T) {
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
	}
	for i, tt := range tests {
		if got, want := intToRoman(tt.i), tt.r; got != want {
			t.Errorf("%d: intToRoman: got %v, want %v", i, got, want)
		}
	}
}
