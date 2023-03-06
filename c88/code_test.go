package leetcode

import (
	"reflect"
	"testing"
)

func TestMerge(t *testing.T) {
	tests := []struct {
		num1    []int
		m       int
		num2    []int
		n       int
		results []int
	}{
		{
			num1:    []int{1, 2, 3, 0, 0, 0},
			m:       3,
			num2:    []int{2, 5, 6},
			n:       3,
			results: []int{1, 2, 2, 3, 5, 6},
		},
		{
			num1:    []int{1},
			m:       1,
			num2:    []int{},
			n:       0,
			results: []int{1},
		},
		{
			num1:    []int{0},
			m:       0,
			num2:    []int{1},
			n:       1,
			results: []int{1},
		},
		{
			num1:    []int{2, 0},
			m:       1,
			num2:    []int{1},
			n:       1,
			results: []int{1, 2},
		},
	}
	for _, tt := range tests {
		merge(tt.num1, tt.m, tt.num2, tt.n)
		if got, want := tt.num1, tt.results; !reflect.DeepEqual(got, want) {
			t.Fatalf("results: got %v, want %v", got, want)
		}
	}
}
