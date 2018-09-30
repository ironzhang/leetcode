package leetcode

import "container/heap"

type HeapImpl struct {
	values []*ListNode
}

func (p *HeapImpl) Len() int {
	return len(p.values)
}

func (p *HeapImpl) Less(i, j int) bool {
	return p.values[i].Val < p.values[j].Val
}

func (p *HeapImpl) Swap(i, j int) {
	p.values[i], p.values[j] = p.values[j], p.values[i]
}

func (p *HeapImpl) Push(x interface{}) {
	v := x.(*ListNode)
	p.values = append(p.values, v)
}

func (p *HeapImpl) Pop() interface{} {
	v := p.values[len(p.values)-1]
	p.values = p.values[:len(p.values)-1]
	return v
}

type PriorityQueue struct {
	HeapImpl
}

func (p *PriorityQueue) Push(x *ListNode) {
	heap.Push(&p.HeapImpl, x)
}

func (p *PriorityQueue) Pop() *ListNode {
	x := heap.Pop(&p.HeapImpl)
	return x.(*ListNode)
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeKLists(lists []*ListNode) *ListNode {
	var queue PriorityQueue
	for _, l := range lists {
		for p := l; p != nil; p = p.Next {
			queue.Push(p)
		}
	}

	h := &ListNode{}
	p := h
	for queue.Len() > 0 {
		p.Next = queue.Pop()
		p = p.Next
	}
	p.Next = nil
	return h.Next
}
