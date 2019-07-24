package leetcode

import "testing"

func TestSearchInsert(t *testing.T) {
	tests := []struct {
		input  []int
		target int
		output int
	}{
		{
			input:  []int{},
			target: 0,
			output: 0,
		},
		{
			input:  []int{1, 3, 5, 6},
			target: 5,
			output: 2,
		},
		{
			input:  []int{1, 3, 5, 6},
			target: 2,
			output: 1,
		},
		{
			input:  []int{1, 3, 5, 6},
			target: 7,
			output: 4,
		},
		{
			input:  []int{1, 3, 5, 6},
			target: 0,
			output: 0,
		},
	}
	for i, tt := range tests {
		n := searchInsert(tt.input, tt.target)
		if got, want := n, tt.output; got != want {
			t.Errorf("%d: got %v, want %v", i, got, want)
		} else {
			t.Logf("%d: got %v", i, got)
		}
	}
}
