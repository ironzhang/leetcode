package leetcode

import "testing"

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

func TestMergeTwoLists(t *testing.T) {
	tests := []struct {
		list1 *ListNode
		list2 *ListNode
		merge *ListNode
	}{
		{
			list1: nil,
			list2: nil,
			merge: nil,
		},
		{
			list1: &ListNode{Val: 0},
			list2: nil,
			merge: &ListNode{Val: 0},
		},
		{
			list1: nil,
			list2: &ListNode{Val: 0},
			merge: &ListNode{Val: 0},
		},
		{
			list1: &ListNode{Val: 0},
			list2: &ListNode{Val: 0},
			merge: &ListNode{Val: 0, Next: &ListNode{Val: 0}},
		},
		{
			list1: &ListNode{Val: 1},
			list2: &ListNode{Val: 0},
			merge: &ListNode{Val: 0, Next: &ListNode{Val: 1}},
		},
		{
			list1: &ListNode{Val: 1},
			list2: &ListNode{Val: 0, Next: &ListNode{Val: 2, Next: &ListNode{Val: 5}}},
			merge: &ListNode{Val: 0, Next: &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 5}}}},
		},
		{
			list1: &ListNode{Val: 1, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4}}},
			list2: &ListNode{Val: 0, Next: &ListNode{Val: 2, Next: &ListNode{Val: 5}}},
			merge: &ListNode{
				Val: 0, Next: &ListNode{
					Val: 1, Next: &ListNode{
						Val: 2, Next: &ListNode{
							Val: 3, Next: &ListNode{
								Val: 4, Next: &ListNode{
									Val: 5, Next: nil,
								},
							},
						},
					},
				},
			},
		},
	}
	for i, tt := range tests {
		if got, want := mergeTwoLists(tt.list1, tt.list2), tt.merge; !listEqual(got, want) {
			t.Errorf("%d: mergeTwoLists: got %v, want %v", i, list2Slice(got), list2Slice(want))
		}
	}
}
