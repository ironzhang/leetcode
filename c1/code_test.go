package leetcode

import (
	"reflect"
	"testing"
)

func TestTwoSum(t *testing.T) {
	tests := []struct {
		nums   []int
		target int
		want   []int
	}{
		{
			nums:   []int{2, 7, 11, 15},
			target: 9,
			want:   []int{0, 1},
		},
		{
			nums:   []int{2, 2, 11, 15},
			target: 4,
			want:   []int{0, 1},
		},
		{
			nums:   []int{2, 2, 11, 15},
			target: 13,
			want:   []int{0, 2},
		},
	}

	for i, test := range tests {
		if got, want := twoSum1(test.nums, test.target), test.want; !reflect.DeepEqual(got, want) {
			t.Errorf("testcase(%d): twoSum1: %v != %v", i, got, want)
		}
		if got, want := twoSum2(test.nums, test.target), test.want; !reflect.DeepEqual(got, want) {
			t.Errorf("testcase(%d): twoSum2: %v != %v", i, got, want)
		}
	}
}

func BenchmarkTwoSum1(b *testing.B) {
	target := 9
	nums := make([]int, 1000)
	nums = append(nums, 2, 7, 11, 15)
	for i := 0; i < b.N; i++ {
		twoSum1(nums, target)
	}
}

func BenchmarkTwoSum2(b *testing.B) {
	target := 9
	nums := make([]int, 1000)
	nums = append(nums, 2, 7, 11, 15)
	for i := 0; i < b.N; i++ {
		twoSum2(nums, target)
	}
}
