package leetcode

import (
	"reflect"
	"testing"
)

func listEqual(h1, h2 *ListNode) bool {
	p1, p2 := h1, h2
	for ; p1 != nil && p2 != nil && p1.Val == p2.Val; p1, p2 = p1.Next, p2.Next {
	}
	if p1 != nil || p2 != nil {
		return false
	}
	return true
}

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

func TestListEqual(t *testing.T) {
	tests := []struct {
		list1 *ListNode
		list2 *ListNode
		equal bool
	}{
		{
			list1: nil,
			list2: nil,
			equal: true,
		},
		{
			list1: nil,
			list2: &ListNode{Val: 0},
			equal: false,
		},
		{
			list1: &ListNode{Val: 0},
			list2: &ListNode{Val: 0},
			equal: true,
		},
		{
			list1: &ListNode{Val: 0, Next: &ListNode{Val: 1}},
			list2: &ListNode{Val: 0, Next: &ListNode{Val: 1}},
			equal: true,
		},
		{
			list1: &ListNode{Val: 0, Next: &ListNode{Val: 1}},
			list2: &ListNode{Val: 0, Next: &ListNode{Val: 1, Next: &ListNode{Val: 2}}},
			equal: false,
		},
		{
			list1: &ListNode{Val: 0, Next: &ListNode{Val: 1, Next: &ListNode{Val: 1}}},
			list2: &ListNode{Val: 0, Next: &ListNode{Val: 1, Next: &ListNode{Val: 2}}},
			equal: false,
		},
		{
			list1: &ListNode{Val: 0, Next: &ListNode{Val: 1, Next: &ListNode{Val: 2}}},
			list2: &ListNode{Val: 0, Next: &ListNode{Val: 1, Next: &ListNode{Val: 2}}},
			equal: true,
		},
	}
	for i, tt := range tests {
		if got, want := listEqual(tt.list1, tt.list2), tt.equal; got != want {
			t.Errorf("%d: ListEqual: got %v, want %v", i, got, want)
		}
	}
}

func TestSlice2List(t *testing.T) {
	tests := []struct {
		slice []int
		list  *ListNode
	}{
		{
			slice: nil,
			list:  nil,
		},
		{
			slice: []int{0},
			list:  &ListNode{Val: 0},
		},
		{
			slice: []int{0, 1},
			list:  &ListNode{Val: 0, Next: &ListNode{Val: 1}},
		},
		{
			slice: []int{0, 1, 2},
			list:  &ListNode{Val: 0, Next: &ListNode{Val: 1, Next: &ListNode{Val: 2}}},
		},
		{
			slice: []int{0, 1, 1},
			list:  &ListNode{Val: 0, Next: &ListNode{Val: 1, Next: &ListNode{Val: 1}}},
		},
	}
	for i, tt := range tests {
		if got, want := slice2List(tt.slice), tt.list; !listEqual(got, want) {
			t.Errorf("%d: slice2List: got %v, want %v", i, list2Slice(got), list2Slice(want))
		}
	}
}

func TestSeek(t *testing.T) {
	list := &ListNode{
		Val: 0,
		Next: &ListNode{
			Val: 1,
			Next: &ListNode{
				Val: 2,
				Next: &ListNode{
					Val: 3,
					Next: &ListNode{
						Val: 4,
						Next: &ListNode{
							Val:  5,
							Next: nil,
						},
					},
				},
			},
		},
	}
	tests := []struct {
		head *ListNode
		n    int
		seek *ListNode
	}{
		{
			head: nil,
			n:    0,
			seek: nil,
		},
		{
			head: nil,
			n:    1,
			seek: nil,
		},
		{
			head: list,
			n:    0,
			seek: list,
		},
		{
			head: list,
			n:    1,
			seek: list.Next,
		},
		{
			head: list,
			n:    2,
			seek: list.Next.Next,
		},
		{
			head: list,
			n:    3,
			seek: list.Next.Next.Next,
		},
		{
			head: list,
			n:    4,
			seek: list.Next.Next.Next.Next,
		},
		{
			head: list,
			n:    5,
			seek: list.Next.Next.Next.Next.Next,
		},
		{
			head: list,
			n:    6,
			seek: list.Next.Next.Next.Next.Next.Next,
		},
		{
			head: list,
			n:    7,
			seek: nil,
		},
	}
	for i, tt := range tests {
		if got, want := seek(tt.head, tt.n), tt.seek; got != want {
			t.Errorf("%d: seek: got %v, want %v", i, list2Slice(got), list2Slice(want))
		}
	}
}

func TestFindNthFromEnd(t *testing.T) {
	list := &ListNode{
		Val: 0,
		Next: &ListNode{
			Val: 1,
			Next: &ListNode{
				Val: 2,
				Next: &ListNode{
					Val: 3,
					Next: &ListNode{
						Val: 4,
						Next: &ListNode{
							Val:  5,
							Next: nil,
						},
					},
				},
			},
		},
	}
	tests := []struct {
		head *ListNode
		n    int
		prev *ListNode
		node *ListNode
	}{
		{
			head: list,
			n:    0,
			prev: list.Next.Next.Next.Next,
			node: list.Next.Next.Next.Next.Next,
		},
		{
			head: list,
			n:    1,
			prev: list.Next.Next.Next,
			node: list.Next.Next.Next.Next,
		},
		{
			head: list,
			n:    2,
			prev: list.Next.Next,
			node: list.Next.Next.Next,
		},
		{
			head: list,
			n:    3,
			prev: list.Next,
			node: list.Next.Next,
		},
		{
			head: list,
			n:    4,
			prev: list,
			node: list.Next,
		},
		{
			head: list,
			n:    5,
			prev: nil,
			node: list,
		},
		{
			head: list,
			n:    6,
			prev: nil,
			node: nil,
		},
		{
			head: list,
			n:    10,
			prev: nil,
			node: nil,
		},
	}
	for i, tt := range tests {
		prev, node := findNthFromEnd(tt.head, tt.n)
		if got, want := prev, tt.prev; got != want {
			t.Errorf("%d: findNthFromEnd: prev: got %v, want %v", i, list2Slice(got), list2Slice(want))
		}
		if got, want := node, tt.node; got != want {
			t.Errorf("%d: findNthFromEnd: node: got %v, want %v", i, list2Slice(got), list2Slice(want))
		}
	}
}

func TestRemoveNthFromEnd(t *testing.T) {
	tests := []struct {
		input  []int
		n      int
		output []int
	}{
		{
			input:  nil,
			n:      0,
			output: nil,
		},
		{
			input:  []int{1},
			n:      0,
			output: []int{1},
		},
		{
			input:  []int{1, 2, 3, 4, 5},
			n:      0,
			output: []int{1, 2, 3, 4, 5},
		},
		{
			input:  []int{1, 2, 3, 4, 5},
			n:      1,
			output: []int{1, 2, 3, 4},
		},
		{
			input:  []int{1, 2, 3, 4, 5},
			n:      2,
			output: []int{1, 2, 3, 5},
		},
	}
	for i, tt := range tests {
		if got, want := list2Slice(removeNthFromEnd(slice2List(tt.input), tt.n)), tt.output; !reflect.DeepEqual(got, want) {
			t.Errorf("%d: removeNthFromEnd: got %v, want %v", i, got, want)
		}
	}
}
