package leetcode

import (
	"reflect"
	"testing"
)

func TestRemoveElement(t *testing.T) {
	tests := []struct {
		nums    []int
		val     int
		results []int
	}{
		{
			nums:    []int{},
			val:     0,
			results: []int{},
		},
		{
			nums:    []int{1},
			val:     1,
			results: []int{},
		},
		{
			nums:    []int{1, 1},
			val:     1,
			results: []int{},
		},
		{
			nums:    []int{3, 2, 2, 3},
			val:     3,
			results: []int{2, 2},
		},
		{
			nums:    []int{0, 1, 2, 2, 3, 0, 4, 2},
			val:     2,
			results: []int{0, 1, 3, 0, 4},
		},
	}
	for i, tt := range tests {
		n := removeElement(tt.nums, tt.val)
		if got, want := tt.nums[:n], tt.results; !reflect.DeepEqual(got, want) {
			t.Fatalf("%d: got %v, want %v", i, got, want)
		} else {
			t.Logf("%d: got %v", i, got)
		}
	}
}
