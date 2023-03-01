package leetcode

import (
	"reflect"
	"testing"
)

func TestLargestLocal(t *testing.T) {
	tests := []struct {
		grid     [][]int
		maxLocal [][]int
	}{
		{
			grid: [][]int{
				{9, 9, 8, 1},
				{5, 6, 2, 6},
				{8, 2, 6, 4},
				{6, 2, 2, 2},
			},
			maxLocal: [][]int{
				{9, 9},
				{8, 6},
			},
		},
		{
			grid: [][]int{
				{1, 1, 1, 1, 1},
				{1, 1, 1, 1, 1},
				{1, 1, 2, 1, 1},
				{1, 1, 1, 1, 1},
				{1, 1, 1, 1, 1},
			},
			maxLocal: [][]int{
				{2, 2, 2},
				{2, 2, 2},
				{2, 2, 2},
			},
		},
	}
	for i, tt := range tests {
		maxLocal := largestLocal(tt.grid)
		if got, want := maxLocal, tt.maxLocal; !reflect.DeepEqual(got, want) {
			t.Fatalf("%d: max local: got %v, want %v", i, got, want)
		} else {
			t.Logf("%d: max local: %v", i, got)
		}
	}
}
