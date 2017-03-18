package leetcode

import (
	"reflect"
	"testing"
)

func ToSlices(n *ListNode) []int {
	values := make([]int, 0)
	for ; n != nil; n = n.Next {
		values = append(values, n.Val)
	}
	return values
}

func TestList(t *testing.T) {
	tests := []struct {
		values []int
	}{
		{values: []int{}},
		{values: []int{1}},
		{values: []int{1, 2}},
		{values: []int{1, 2, 3}},
		{values: []int{1, 1, 1, 1, 1}},
	}

	for i, test := range tests {
		var list List
		for _, v := range test.values {
			list.AddNode(&ListNode{Val: v})
		}
		if got, want := ToSlices(list.Head), test.values; !reflect.DeepEqual(got, want) {
			t.Errorf("testcase(%d): %v != %v", i, got, want)
		}
	}
}

func TestAddTwoNumbers(t *testing.T) {
	tests := []struct {
		input1 []int
		input2 []int
		output []int
	}{
		{
			input1: []int{},
			input2: []int{},
			output: []int{},
		},
		{
			input1: []int{2, 4, 3},
			input2: []int{5, 6, 4},
			output: []int{7, 0, 8},
		},
		{
			input1: []int{},
			input2: []int{5, 6, 4},
			output: []int{5, 6, 4},
		},
		{
			input1: []int{2, 4, 3},
			input2: []int{},
			output: []int{2, 4, 3},
		},
		{
			input1: []int{2, 4, 5},
			input2: []int{5, 6, 4},
			output: []int{7, 0, 0, 1},
		},
		{
			input1: []int{2, 4},
			input2: []int{5, 6, 4},
			output: []int{7, 0, 5},
		},
		{
			input1: []int{2, 4, 9},
			input2: []int{5, 6},
			output: []int{7, 0, 0, 1},
		},
		{
			input1: []int{1},
			input2: []int{1, 1, 0},
			output: []int{2, 1, 0},
		},
	}

	for i, test := range tests {
		var l1, l2 List
		for _, v := range test.input1 {
			l1.AddNode(&ListNode{Val: v})
		}
		for _, v := range test.input2 {
			l2.AddNode(&ListNode{Val: v})
		}
		l := addTwoNumbers(l1.Head, l2.Head)
		if got, want := ToSlices(l), test.output; !reflect.DeepEqual(got, want) {
			t.Errorf("testcase(%d): %v != %v", i, got, want)
		}
	}
}
