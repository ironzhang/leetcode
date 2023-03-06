package leetcode

import (
	"reflect"
	"testing"
)

func TestIntersect(t *testing.T) {
	tests := []struct {
		nums1   []int
		nums2   []int
		results []int
	}{
		{
			nums1:   []int{1, 2, 2, 1},
			nums2:   []int{2, 2},
			results: []int{2, 2},
		},
		{
			nums1:   []int{4, 9, 5},
			nums2:   []int{9, 4, 9, 8, 4},
			results: []int{4, 9},
		},
	}
	for _, tt := range tests {
		if got, want := intersect(tt.nums1, tt.nums2), tt.results; !reflect.DeepEqual(got, want) {
			t.Fatalf("results: got %v, want %v", got, want)
		}
	}
}
