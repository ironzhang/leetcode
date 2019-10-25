package leetcode

import (
	"container/list"
	"fmt"
	"sort"
)

type Node struct {
	target     int
	candidates []int
}

func (n *Node) LastCandidate() int {
	if x := len(n.candidates); x > 0 {
		return n.candidates[x-1]
	}
	return 0
}

type Queue struct {
	list list.List
}

func (q *Queue) Push(n Node) {
	q.list.PushBack(n)
}

func (q *Queue) Pop() (Node, bool) {
	e := q.list.Front()
	if e == nil {
		return Node{}, false
	}
	q.list.Remove(e)
	return e.Value.(Node), true
}

type Answer struct {
	m map[string]struct{}
	a [][]int
}

func (p *Answer) Add(a []int) {
	sort.Ints(a)
	k := fmt.Sprintf("%v", a)
	if _, ok := p.m[k]; ok {
		return
	}
	p.m[k] = struct{}{}
	p.a = append(p.a, a)
}

func Append(src []int, x int) []int {
	dst := make([]int, len(src)+1)
	copy(dst, src)
	dst[len(src)] = x
	return dst
}

func combinationSum(candidates []int, target int) [][]int {
	return bfs1(candidates, target)
}

func bfs0(candidates []int, target int) [][]int {
	ans := Answer{m: make(map[string]struct{})}
	var q Queue
	q.Push(Node{target: target})
	for {
		n, ok := q.Pop()
		if !ok {
			break
		}
		for _, c := range candidates {
			x := n.target - c
			if x == 0 {
				ans.Add(Append(n.candidates, c))
			} else if x > 0 {
				nn := Node{target: x, candidates: Append(n.candidates, c)}
				q.Push(nn)
			}
		}
	}
	return ans.a
}

func bfs1(candidates []int, target int) (ans [][]int) {
	sort.Ints(candidates)
	var q Queue
	q.Push(Node{target: target})
	for {
		n, ok := q.Pop()
		if !ok {
			break
		}
		for _, c := range candidates {
			if n.LastCandidate() > c {
				continue
			}
			x := n.target - c
			if x == 0 {
				ans = append(ans, Append(n.candidates, c))
			} else if x > 0 {
				nn := Node{target: x, candidates: Append(n.candidates, c)}
				q.Push(nn)
			} else {
				break
			}
		}
	}
	return ans
}
