package leetcode

import (
	"math"
	"testing"
)

func TestMyAtoi(t *testing.T) {
	tests := []struct {
		str string
		val int
	}{
		{"10", 10},
		{"-10", -10},
		{"  -10", -10},
		{"  10", 10},
		{"  a10", 0},
		{"10a", 10},
		{"1a0", 1},
		{"a10", 0},
		{"", 0},
		{"2147483648", math.MaxInt32},
		{"-2147483649", math.MinInt32},
	}
	for _, tt := range tests {
		if got, want := myAtoi(tt.str), tt.val; got != want {
			t.Errorf("my atoi: str[%s], got[%d], want[%d]", tt.str, got, want)
		}
	}
}
