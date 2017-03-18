package leetcode

import "testing"

func TestFindMedianSortedArrays(t *testing.T) {
	tests := []struct {
		nums1 []int
		nums2 []int
		want  float64
	}{
		{
			nums1: []int{1, 3},
			nums2: []int{2},
			want:  2,
		},
		{
			nums1: []int{1, 2},
			nums2: []int{3, 4},
			want:  2.5,
		},
	}

	for i, test := range tests {
		got, want := findMedianSortedArrays(test.nums1, test.nums2), test.want
		if got != want {
			t.Errorf("testcase(%d): %v != %v", i, got, want)
		}
	}
}
