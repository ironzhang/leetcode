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

func sort(a []int) []int {
	var q PriorityQueue
	for _, x := range a {
		q.Push(&ListNode{Val: x})
	}

	b := make([]int, 0, len(a))
	for q.Len() > 0 {
		n := q.Pop()
		b = append(b, n.Val)
	}
	return b
}

func TestPriorityQueue(t *testing.T) {
	tests := []struct {
		input  []int
		output []int
	}{
		{
			input:  []int{1},
			output: []int{1},
		},
		{
			input:  []int{2, 1},
			output: []int{1, 2},
		},
		{
			input:  []int{2, 1, 3},
			output: []int{1, 2, 3},
		},
	}
	for i, tt := range tests {
		if got, want := sort(tt.input), tt.output; !reflect.DeepEqual(got, want) {
			t.Errorf("%d: sort: got %v, want %v", i, got, want)
		}
	}
}

func TestMergeKLists(t *testing.T) {
	tests := []struct {
		lists []*ListNode
		merge *ListNode
	}{
		{
			lists: []*ListNode{
				&ListNode{Val: 1, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5}}},
				&ListNode{Val: 1, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4}}},
				&ListNode{Val: 2, Next: &ListNode{Val: 6}},
			},
			merge: &ListNode{
				Val: 1, Next: &ListNode{
					Val: 1, Next: &ListNode{
						Val: 2, Next: &ListNode{
							Val: 3, Next: &ListNode{
								Val: 4, Next: &ListNode{
									Val: 4, Next: &ListNode{
										Val: 5, Next: &ListNode{
											Val: 6, Next: nil,
										},
									},
								},
							},
						},
					},
				},
			},
		},
		{
			lists: []*ListNode{
				&ListNode{Val: -2, Next: &ListNode{Val: -1, Next: &ListNode{Val: -1, Next: &ListNode{Val: -1}}}},
				nil,
			},
			merge: &ListNode{Val: -2, Next: &ListNode{Val: -1, Next: &ListNode{Val: -1, Next: &ListNode{Val: -1}}}},
		},
	}
	for i, tt := range tests {
		if got, want := mergeKLists(tt.lists), tt.merge; !listEqual(got, want) {
			t.Errorf("%d: mergeKLists: got %v, want %v", i, list2Slice(got), list2Slice(want))
		}
	}
}
