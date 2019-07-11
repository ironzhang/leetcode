package leetcode

import (
	"reflect"
	"testing"
)

func TestNextPermutation(t *testing.T) {
	tests := []struct {
		nums []int
		res  []int
	}{
		{
			nums: []int{5, 1, 1},
			res:  []int{1, 1, 5},
		},
		{
			nums: []int{1, 1},
			res:  []int{1, 1},
		},
		{
			nums: []int{1, 5, 1},
			res:  []int{5, 1, 1},
		},
		{
			nums: []int{1, 2, 3},
			res:  []int{1, 3, 2},
		},
		{
			nums: []int{3, 2, 1},
			res:  []int{1, 2, 3},
		},
		{
			nums: []int{1, 1, 5},
			res:  []int{1, 5, 1},
		},
	}
	for i, tt := range tests {
		nextPermutation(tt.nums)
		if got, want := tt.nums, tt.res; !reflect.DeepEqual(got, want) {
			t.Errorf("%d: got %v, want %v", i, got, want)
		} else {
			t.Logf("%d: got %v", i, got)
		}
	}
}

func TestReverse(t *testing.T) {
	tests := []struct {
		input  []int
		output []int
	}{
		{
			input:  []int{1, 2, 3},
			output: []int{3, 2, 1},
		},
		{
			input:  []int{1, 2, 3, 4},
			output: []int{4, 3, 2, 1},
		},
	}
	for i, tt := range tests {
		reverse(tt.input)
		if got, want := tt.input, tt.output; !reflect.DeepEqual(got, want) {
			t.Errorf("%d: got %v, want %v", i, got, want)
		} else {
			t.Logf("%d: got %v", i, got)
		}
	}
}
