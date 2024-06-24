package leetcode

import "testing"

func TestTrap(t *testing.T) {
	tests := []struct {
		height []int
		ans    int
	}{
		{
			height: []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1},
			ans:    6,
		},
		{
			height: []int{4, 2, 0, 3, 2, 5},
			ans:    9,
		},
	}
	for i, tt := range tests {
		ans := trap(tt.height)
		if got, want := ans, tt.ans; got != want {
			t.Errorf("testcase(%d): trap: %v != %v", i, got, want)
		}
	}
}
