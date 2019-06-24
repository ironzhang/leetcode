package leetcode

import "testing"

func TestIsPalindrome(t *testing.T) {
	tests := []struct {
		grid   [][]byte
		number int
	}{
		{
			grid: [][]byte{
				{'1', '1', '1', '1', '0'},
				{'1', '1', '0', '1', '0'},
				{'1', '1', '0', '0', '0'},
				{'0', '0', '0', '0', '0'},
			},
			number: 1,
		},
		{
			grid: [][]byte{
				{'1', '1', '0', '0', '0'},
				{'1', '1', '0', '0', '0'},
				{'0', '0', '1', '0', '0'},
				{'0', '0', '0', '1', '1'},
			},
			number: 3,
		},
	}
	for i, tt := range tests {
		if got, want := numIslands(tt.grid), tt.number; got != want {
			t.Errorf("%d: numIslands: got=%v, want=%v", i, got, want)
		} else {
			t.Logf("%d: numIslands: got=%v", i, got)
		}
	}
}
