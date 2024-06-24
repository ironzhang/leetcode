package leetcode

import "testing"

func TestUniquePaths(t *testing.T) {
	tests := []struct {
		m      int
		n      int
		result int
	}{
		{
			m:      3,
			n:      7,
			result: 28,
		},
		{
			m:      3,
			n:      2,
			result: 3,
		},
		{
			m:      3,
			n:      3,
			result: 6,
		},
	}
	for i, tt := range tests {
		result := uniquePaths(tt.m, tt.n)
		if got, want := result, tt.result; got != want {
			t.Errorf("testcase(%d): uniquePaths: %v != %v", i, got, want)
		}
	}
}
