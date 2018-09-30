package leetcode

import (
	"reflect"
	"testing"
)

func slice2List(slice []int) *ListNode {
	var head, prev *ListNode
	for _, val := range slice {
		node := &ListNode{Val: val}
		if head == nil {
			head = node
		}
		if prev != nil {
			prev.Next = node
		}
		prev = node
	}
	return head
}

func list2Slice(head *ListNode) []int {
	var slice []int
	for p := head; p != nil; p = p.Next {
		slice = append(slice, p.Val)
	}
	return slice
}

func TestSwapPairs(t *testing.T) {
	tests := []struct {
		input  []int
		output []int
	}{
		{
			input:  nil,
			output: nil,
		},
		{
			input:  []int{1},
			output: []int{1},
		},
		{
			input:  []int{1, 2},
			output: []int{2, 1},
		},
		{
			input:  []int{1, 2, 3},
			output: []int{2, 1, 3},
		},
		{
			input:  []int{1, 2, 3, 4},
			output: []int{2, 1, 4, 3},
		},
	}
	for i, tt := range tests {
		input := slice2List(tt.input)
		ouput := swapPairs(input)
		if got, want := list2Slice(ouput), tt.output; !reflect.DeepEqual(got, want) {
			t.Errorf("%d: swapPairs: got %v, want %v", i, got, want)
		}
	}
}
