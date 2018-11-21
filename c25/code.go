package leetcode

import "github.com/ironzhang/leetcode/util"

type ListNode = util.ListNode

type stack []*ListNode

func (p *stack) Push(n *ListNode) {
	*p = append(*p, n)
}

func (p *stack) Pop() (*ListNode, bool) {
	if l := len(*p); l > 0 {
		n := (*p)[l-1]
		*p = (*p)[:l-1]
		return n, true
	}
	return nil, false
}

func reverseStack(s *stack) (start, end *ListNode) {
	var prev *ListNode
	for {
		n, ok := s.Pop()
		if !ok {
			break
		}
		prev = n
	}
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	s := make(stack, 0, k)
	for p := head; p != nil; p = p.Next {
		s.Push(p)
		if len(s) == k {

		}
	}
}
