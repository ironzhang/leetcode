package leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

type List struct {
	Head *ListNode
	Tail *ListNode
}

func (l *List) AddNode(n *ListNode) {
	if l.Head == nil {
		l.Head = n
	}
	if l.Tail != nil {
		l.Tail.Next = n
	}
	l.Tail = n
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var list List
	var sum, carry, val int

	n1, n2 := l1, l2
	for n1 != nil || n2 != nil || carry > 0 {
		sum = carry
		if n1 != nil {
			sum += n1.Val
		}
		if n2 != nil {
			sum += n2.Val
		}

		val, carry = sum%10, sum/10
		list.AddNode(&ListNode{Val: val})

		if n1 != nil {
			n1 = n1.Next
		}
		if n2 != nil {
			n2 = n2.Next
		}
	}

	return list.Head
}
