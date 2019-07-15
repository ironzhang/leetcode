package leetcode

import (
	"testing"
)

func TestSearchRotatedIndex(t *testing.T) {
	tests := []struct {
		nums  []int
		index int
	}{
		{
			nums:  []int{4, 5, 6, 7, 0, 1, 2},
			index: 4,
		},
		{
			nums:  []int{5, 6, 7, 0, 1, 2},
			index: 3,
		},
		{
			nums:  []int{5, 1, 2, 3},
			index: 1,
		},
		{
			nums:  []int{1, 2, 3},
			index: 3,
		},
		{
			nums:  []int{2, 3, 1},
			index: 2,
		},
		{
			nums:  []int{1, 2},
			index: 2,
		},
		{
			nums:  []int{2, 1},
			index: 1,
		},
	}
	for i, tt := range tests {
		index := searchRotatedIndex(tt.nums)
		if got, want := index, tt.index; got != want {
			t.Errorf("%d: got %v, want %v", i, got, want)
		} else {
			t.Logf("%d: got %v", i, got)
		}
	}
}

func TestBinSearch(t *testing.T) {
	tests := []struct {
		nums   []int
		target int
		output int
	}{
		{
			nums:   []int{0, 1, 2, 4, 5, 6, 7},
			target: 0,
			output: 0,
		},
		{
			nums:   []int{0, 1, 2, 4, 5, 6, 7},
			target: 3,
			output: -1,
		},
	}
	for i, tt := range tests {
		index := binSearch(tt.nums, 0, len(tt.nums), tt.target)
		if got, want := index, tt.output; got != want {
			t.Errorf("%d: got %v, want %v", i, got, want)
		} else {
			t.Logf("%d: got %v", i, got)
		}
	}
}

func TestSearch(t *testing.T) {
	tests := []struct {
		nums   []int
		target int
		output int
	}{
		{
			nums:   []int{},
			target: 5,
			output: -1,
		},
		{
			nums:   []int{1},
			target: 1,
			output: 0,
		},
		{
			nums:   []int{1, 3},
			target: 3,
			output: 1,
		},
		{
			nums:   []int{4, 5, 6, 7, 0, 1, 2},
			target: 0,
			output: 4,
		},
		{
			nums:   []int{4, 5, 6, 7, 0, 1, 2},
			target: 3,
			output: -1,
		},
	}
	for i, tt := range tests {
		index := search(tt.nums, tt.target)
		if got, want := index, tt.output; got != want {
			t.Errorf("%d: got %v, want %v", i, got, want)
		} else {
			t.Logf("%d: got %v", i, got)
		}
	}
}
