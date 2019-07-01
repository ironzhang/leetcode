package leetcode

import (
	"reflect"
	"testing"

	"github.com/ironzhang/leetcode/util"
)

func TestReverseKNodes(t *testing.T) {
	tests := []struct {
		input  []int
		k      int
		output []int
		prev   []int
		next   []int
	}{
		{
			input:  []int{1, 2, 3},
			k:      2,
			output: []int{2, 1, 3},
			prev:   []int{1, 3},
			next:   []int{3},
		},
		{
			input:  []int{1, 2, 3, 4, 5},
			k:      2,
			output: []int{2, 1, 3, 4, 5},
			prev:   []int{1, 3, 4, 5},
			next:   []int{3, 4, 5},
		},
		{
			input:  []int{1, 2, 3, 4, 5},
			k:      3,
			output: []int{3, 2, 1, 4, 5},
			prev:   []int{1, 4, 5},
			next:   []int{4, 5},
		},
	}
	for i, tt := range tests {
		var n ListNode
		h := util.Slice2List(tt.input)
		prev, next := reverseKNodes(&n, h, tt.k)
		if got, want := util.List2Slice(n.Next), tt.output; !reflect.DeepEqual(got, want) {
			t.Errorf("%d: reverseKNodes: list: got %v, want %v", i, got, want)
		} else {
			t.Logf("%d: reverseKNodes: list: got %v", i, got)
		}
		if got, want := util.List2Slice(prev), tt.prev; !reflect.DeepEqual(got, want) {
			t.Errorf("%d: reverseKNodes: prev: got %v, want %v", i, got, want)
		} else {
			t.Logf("%d: reverseKNodes: prve: got %v", i, got)
		}
		if got, want := util.List2Slice(next), tt.next; !reflect.DeepEqual(got, want) {
			t.Errorf("%d: reverseKNodes: next: got %v, want %v", i, got, want)
		} else {
			t.Logf("%d: reverseKNodes: next: got %v", i, got)
		}
	}
}

func TestReverseKGroup(t *testing.T) {
	tests := []struct {
		input  []int
		k      int
		output []int
	}{
		{
			input:  []int{1, 2, 3},
			k:      2,
			output: []int{2, 1, 3},
		},
		{
			input:  []int{1, 2, 3, 4, 5},
			k:      2,
			output: []int{2, 1, 4, 3, 5},
		},
		{
			input:  []int{1, 2, 3, 4, 5},
			k:      3,
			output: []int{3, 2, 1, 4, 5},
		},
		{
			input:  []int{1, 2, 3, 4, 5},
			k:      1,
			output: []int{1, 2, 3, 4, 5},
		},
		{
			input:  []int{1, 2, 3, 4, 5},
			k:      5,
			output: []int{5, 4, 3, 2, 1},
		},
		{
			input:  []int{1, 2, 3, 4, 5},
			k:      4,
			output: []int{4, 3, 2, 1, 5},
		},
	}
	for i, tt := range tests {
		input := util.Slice2List(tt.input)
		output := reverseKGroup(input, tt.k)
		if got, want := util.List2Slice(output), tt.output; !reflect.DeepEqual(got, want) {
			t.Errorf("%d: reverseKGroup: got %v, want %v", i, got, want)
		} else {
			t.Logf("%d: reverseKGroup: got %v", i, got)
		}
	}
}
