package leetcode

import (
	"errors"
)

func isMatch(s string, p string) bool {
	m, err := re2machine(p)
	if err != nil {
		panic(err)
	}
	return m.match(s)
}

type edge struct {
	letter rune
	state  *state
}

type state struct {
	edges []edge
}

type machine struct {
	start *state
	end   *state
}

func re2machine(p string) (*machine, error) {
	var m, nm *machine
	m = newLetterMachine('*')
	for i := 0; i < len(p); i++ {
		if c := rune(p[i]); (c >= 'a' && c <= 'z') || c == '.' {
			if i+1 < len(p) && p[i+1] == '*' {
				nm = newLetterStarMachine(c)
				i++
			} else {
				nm = newLetterMachine(c)
			}
		} else {
			return nil, errors.New("invalid pattern")
		}
		m = merge(m, nm)
	}
	return m, nil
}

func newLetterMachine(c rune) *machine {
	start, end := &state{}, &state{}
	start.edges = append(start.edges, edge{letter: c, state: end})
	return &machine{start: start, end: end}
}

func newLetterStarMachine(c rune) *machine {
	m := newLetterMachine(c)
	si, sf := &state{}, &state{}
	si.edges = append(si.edges, edge{letter: '*', state: m.start})
	si.edges = append(si.edges, edge{letter: '*', state: sf})
	m.end.edges = append(m.end.edges, edge{letter: '*', state: m.start})
	m.end.edges = append(m.end.edges, edge{letter: '*', state: sf})
	return &machine{start: si, end: sf}
}

func merge(m1, m2 *machine) *machine {
	if m1 == nil {
		return m2
	}
	m1.end.edges = append(m1.end.edges, edge{letter: '*', state: m2.start})
	return &machine{start: m1.start, end: m2.end}
}

func (m *machine) match(s string) bool {
	start, end := m.start, m.end
	states := map[*state]bool{start: true}

	states = eclosure(states)
	for _, c := range s {
		states = eclosure(move(states, c))
	}
	if states[end] {
		return true
	}
	return false
}

func eclosure(oldStates map[*state]bool) (newStates map[*state]bool) {
	newStates = make(map[*state]bool)
	for s := range oldStates {
		add(newStates, s)
	}
	return newStates
}

func add(states map[*state]bool, s *state) {
	if !states[s] {
		states[s] = true
		for _, e := range s.edges {
			if e.letter == '*' {
				add(states, e.state)
			}
		}
	}
}

func move(oldStates map[*state]bool, c rune) (newStates map[*state]bool) {
	newStates = make(map[*state]bool)
	for s := range oldStates {
		for _, e := range s.edges {
			if (e.letter == c || e.letter == '.') && !newStates[e.state] {
				newStates[e.state] = true
			}
		}
	}
	return newStates
}

func (m machine) String() string {
	//	var b bytes.Buffer
	//	for _, s := range m {
	//		fmt.Fprintf(&b, "%d: ", s.id)
	//		for _, e := range s.edges {
	//			fmt.Fprintf(&b, "%c:%d, ", e.letter, e.state.id)
	//		}
	//		fmt.Fprintf(&b, "\n")
	//	}
	//	return b.String()
	return ""
}
