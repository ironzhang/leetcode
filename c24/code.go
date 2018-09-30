package leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

func swapPairs(head *ListNode) *ListNode {
	list := &ListNode{Next: head}
	for prev, node := list, list.Next; node != nil; node = node.Next {
		if node.Next != nil {
			swapNode(prev, node, node.Next)
			prev = node
		}
	}
	return list.Next
}

func swapNode(prev, node, next *ListNode) {
	prev.Next = next
	node.Next = next.Next
	next.Next = node
}
