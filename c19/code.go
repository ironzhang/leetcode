package leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	prev, node := findNthFromEnd(head, n-1)
	if node == nil {
		return head
	}
	if prev == nil {
		return head.Next
	}
	prev.Next = node.Next
	return head
}

func findNthFromEnd(head *ListNode, n int) (prev, node *ListNode) {
	p := seek(head, n)
	if p == nil {
		return nil, nil
	}
	for ; p != nil; p = p.Next {
		prev = node
		if node == nil {
			node = head
		} else {
			node = node.Next
		}
	}
	return prev, node
}

func seek(p *ListNode, n int) *ListNode {
	if n < 0 {
		return nil
	}
	for i := 0; i < n; i++ {
		if p == nil {
			return nil
		}
		p = p.Next
	}
	return p
}
