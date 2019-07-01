package util

type ListNode struct {
	Val  int
	Next *ListNode
}

func Slice2List(slice []int) *ListNode {
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

func List2Slice(head *ListNode) []int {
	var slice []int
	for p := head; p != nil; p = p.Next {
		slice = append(slice, p.Val)
	}
	return slice
}
