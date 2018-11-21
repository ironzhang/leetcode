package leetcode

import (
	"reflect"
	"testing"
)

func TestRemoveDuplicates(t *testing.T) {
	tests := []struct {
		nums    []int
		results []int
	}{
		{
			nums:    []int{},
			results: []int{},
		},
		{
			nums:    []int{1},
			results: []int{1},
		},
		{
			nums:    []int{1, 1},
			results: []int{1},
		},
		{
			nums:    []int{1, 1, 2},
			results: []int{1, 2},
		},
		{
			nums:    []int{1, 1, 2, 2},
			results: []int{1, 2},
		},
		{
			nums:    []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4},
			results: []int{0, 1, 2, 3, 4},
		},
	}
	for i, tt := range tests {
		n := removeDuplicates(tt.nums)
		if got, want := tt.nums[:n], tt.results; !reflect.DeepEqual(got, want) {
			t.Fatalf("%d: got %v, want %v", i, got, want)
		} else {
			t.Logf("%d: got %v", i, got)
		}
	}
}
