package leetcode

import (
	"reflect"
	"testing"
)

func TestCombinationSum(t *testing.T) {
	tests := []struct {
		candidates []int
		target     int
		answer     [][]int
	}{
		{
			candidates: []int{2, 3, 6, 7},
			target:     7,
			answer: [][]int{
				{7},
				{2, 2, 3},
			},
		},
		{
			candidates: []int{2, 3, 5},
			target:     8,
			answer: [][]int{
				{3, 5},
				{2, 3, 3},
				{2, 2, 2, 2},
			},
		},
		{
			candidates: []int{7, 3, 2},
			target:     18,
			answer: [][]int{
				{2, 2, 7, 7},
				{2, 3, 3, 3, 7},
				{2, 2, 2, 2, 3, 7},
				{3, 3, 3, 3, 3, 3},
				{2, 2, 2, 3, 3, 3, 3},
				{2, 2, 2, 2, 2, 2, 3, 3},
				{2, 2, 2, 2, 2, 2, 2, 2, 2},
			},
		},
	}

	for i, test := range tests {
		if got, want := combinationSum(test.candidates, test.target), test.answer; !reflect.DeepEqual(got, want) {
			t.Errorf("testcase(%d): combinationSum: (%d)%v != (%d)%v", i, len(got), got, len(want), want)
		}
	}
}
