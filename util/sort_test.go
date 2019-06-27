package util

import (
	"reflect"
	"sort"
	"testing"
)

func TestQuickSort(t *testing.T) {
	tests := [][]int{
		{6, 1, 2, 7, 9, 3, 4, 5, 10, 8},
		{6, 6, 1, 2, 7, 9, 3, 4, 5, 10, 8},
	}

	for _, tt := range tests {
		a := make([]int, len(tt))
		b := make([]int, len(tt))
		copy(a, tt)
		copy(b, tt)
		QuickSort(a)
		sort.Ints(b)
		if got, want := a, b; !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want %v", got, want)
		} else {
			t.Logf("got %v", got)
		}
	}
}
