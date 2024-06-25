package leetcode

import (
	"reflect"
	"testing"
)

func TestPlusOne(t *testing.T) {
	tests := []struct {
		digits []int
		ans    []int
	}{
		{
			digits: []int{1, 2, 3},
			ans:    []int{1, 2, 4},
		},
		{
			digits: []int{4, 3, 2, 1},
			ans:    []int{4, 3, 2, 2},
		},
		{
			digits: []int{0},
			ans:    []int{1},
		},
		{
			digits: []int{9},
			ans:    []int{1, 0},
		},
		{
			digits: []int{9, 9},
			ans:    []int{1, 0, 0},
		},
	}
	for i, tt := range tests {
		ans := plusOne(tt.digits)
		if got, want := ans, tt.ans; !reflect.DeepEqual(got, want) {
			t.Errorf("testcase(%d): plusOne: %v != %v", i, got, want)
		}
	}
}
