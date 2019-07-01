package leetcode

import (
	"github.com/ironzhang/leetcode/util"
)

type ListNode = util.ListNode

func reverseKNodes(prev, head *ListNode, k int) (*ListNode, *ListNode) {
	p := head
	n := head.Next
	for i := 1; i < k && n != nil; i++ {
		next := n.Next
		n.Next = p
		p, n = n, next
	}
	prev.Next = p
	head.Next = n
	return head, head.Next
}

func isKGroup(head *ListNode, k int) bool {
	n := 0
	for p := head; p != nil; p = p.Next {
		n++
		if n >= k {
			return true
		}
	}
	return false
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	node := ListNode{Next: head}
	prev := &node
	for isKGroup(head, k) {
		prev, head = reverseKNodes(prev, head, k)
	}
	return node.Next
}
