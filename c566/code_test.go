package leetcode

import (
	"reflect"
	"testing"
)

func TestMatrixReshape(t *testing.T) {
	tests := []struct {
		mat    [][]int
		r      int
		c      int
		result [][]int
	}{
		{
			mat: [][]int{
				{1, 2},
				{3, 4},
			},
			r: 1,
			c: 4,
			result: [][]int{
				{1, 2, 3, 4},
			},
		},
		{
			mat: [][]int{
				{1, 2},
				{3, 4},
			},
			r: 2,
			c: 4,
			result: [][]int{
				{1, 2},
				{3, 4},
			},
		},
	}
	for _, tt := range tests {
		if got, want := matrixReshape(tt.mat, tt.r, tt.c), tt.result; !reflect.DeepEqual(got, want) {
			t.Fatalf("matrix reshape: got %v, want %v", got, want)
		}
	}
}
