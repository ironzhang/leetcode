package leetcode

import (
	"bytes"
	"errors"
	"fmt"
)

func isMatch(s string, p string) bool {
	g, err := re2graph(p)
	if err != nil {
		panic(err)
	}
	return g.match(s)
}

type edge struct {
	letter rune
	node   *node
}

type node struct {
	edges []edge
}

func newNode() *node {
	return &node{}
}

func (p *node) addEdge(r rune, n *node) {
	p.edges = append(p.edges, edge{letter: r, node: n})
}

type graph struct {
	start, end *node
}

func newGraph(r rune) *graph {
	s, e := newNode(), newNode()
	s.addEdge(r, e)
	return &graph{start: s, end: e}
}

func newStarGraph(r rune) *graph {
	s, e, i, f := newNode(), newNode(), newNode(), newNode()
	s.addEdge(r, e)
	i.addEdge('*', s)
	i.addEdge('*', f)
	e.addEdge('*', s)
	e.addEdge('*', f)
	return &graph{start: i, end: f}
}

func (p *graph) merge(g *graph) {
	p.end.addEdge('*', g.start)
	p.end = g.end
}

func re2graph(p string) (*graph, error) {
	g := newGraph('*')
	for i := 0; i < len(p); i++ {
		if c := rune(p[i]); (c >= 'a' && c <= 'z') || c == '.' {
			if i+1 < len(p) && p[i+1] == '*' {
				g.merge(newStarGraph(c))
				i++
			} else {
				g.merge(newGraph(c))
			}
		} else {
			return nil, errors.New("invalid pattern")
		}
	}
	return g, nil
}

type nodes map[*node]bool

func (p nodes) add(n *node) {
	if !p[n] {
		p[n] = true
		for _, e := range n.edges {
			if e.letter == '*' {
				p.add(e.node)
			}
		}
	}
}

func move(olds nodes, r rune) (news nodes) {
	news = make(nodes)
	for n := range olds {
		for _, e := range n.edges {
			if e.letter == r || e.letter == '.' {
				news.add(e.node)
			}
		}
	}
	return news
}

func (p *graph) match(s string) bool {
	m := make(nodes)
	m.add(p.start)
	for _, c := range s {
		m = move(m, c)
	}
	if m[p.end] {
		return true
	}
	return false
}

func (p *graph) String() string {
	var b bytes.Buffer
	visited := make(map[*node]bool)
	list := p.visit(nil, p.start, visited)

	ids := make(map[*node]int)
	for i, n := range list {
		ids[n] = i
	}
	for i, n := range list {
		fmt.Fprintf(&b, "%d: ", i)
		for _, e := range n.edges {
			fmt.Fprintf(&b, "%c:%d, ", e.letter, ids[e.node])
		}
		fmt.Fprintln(&b)
	}
	return b.String()
}

func (p *graph) visit(list []*node, n *node, visited map[*node]bool) []*node {
	if !visited[n] {
		list = append(list, n)
		visited[n] = true
		for _, e := range n.edges {
			list = p.visit(list, e.node, visited)
		}
	}
	return list
}
