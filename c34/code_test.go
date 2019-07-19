package leetcode

import (
	"reflect"
	"testing"
)

func TestSearchRange(t *testing.T) {
	tests := []struct {
		nums   []int
		target int
		output []int
	}{
		{
			nums:   []int{5, 7, 7, 8, 8, 10},
			target: 8,
			output: []int{3, 4},
		},
		{
			nums:   []int{5, 7, 7, 8, 8, 10},
			target: 6,
			output: []int{-1, -1},
		},
	}
	for i, tt := range tests {
		output := searchRange(tt.nums, tt.target)
		if got, want := output, tt.output; !reflect.DeepEqual(got, want) {
			t.Errorf("%d: got %v, want %v", i, got, want)
		} else {
			t.Logf("%d: got %v", i, got)
		}
	}
}

func TestLeftRangeIndex(t *testing.T) {
	tests := []struct {
		nums   []int
		left   int
		right  int
		target int
		output int
	}{
		{
			nums:   []int{5, 7, 7, 8, 8, 10},
			left:   2,
			right:  4,
			target: 8,
			output: 3,
		},
		{
			nums:   []int{5, 7, 7, 8, 8, 10},
			left:   1,
			right:  4,
			target: 8,
			output: 3,
		},
		{
			nums:   []int{5, 7, 7, 8, 8, 10},
			left:   1,
			right:  3,
			target: 7,
			output: 1,
		},
		{
			nums:   []int{5, 7, 7, 8, 8, 10},
			left:   0,
			right:  1,
			target: 7,
			output: 1,
		},
	}
	for i, tt := range tests {
		output := leftRangeIndex(tt.nums, tt.left, tt.right, tt.target)
		if got, want := output, tt.output; got != want {
			t.Errorf("%d: got %v, want %v", i, got, want)
		} else {
			t.Logf("%d: got %v", i, got)
		}
	}
}

func TestRightRangeIndex(t *testing.T) {
	tests := []struct {
		nums   []int
		left   int
		right  int
		target int
		output int
	}{
		{
			nums:   []int{5, 7, 7, 8, 8, 10},
			left:   3,
			right:  5,
			target: 8,
			output: 4,
		},
		{
			nums:   []int{5, 7, 7, 8, 8, 10},
			left:   4,
			right:  4,
			target: 8,
			output: 4,
		},
		{
			nums:   []int{5, 7, 7, 8, 8, 10},
			left:   1,
			right:  3,
			target: 7,
			output: 2,
		},
	}
	for i, tt := range tests {
		output := rightRangeIndex(tt.nums, tt.left, tt.right, tt.target)
		if got, want := output, tt.output; got != want {
			t.Errorf("%d: got %v, want %v", i, got, want)
		} else {
			t.Logf("%d: got %v", i, got)
		}
	}
}
