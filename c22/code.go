package leetcode

var caches = map[int][]string{
	0: []string{},
	1: []string{"()"},
}

func generateParenthesis(n int) []string {
	if set, ok := caches[n]; ok {
		return set
	}

	set := generateParenthesis(n - 1)
	nset := make([]string, 0, len(set)*3)
	already := make(map[string]bool)
	add := func(s string) {
		if !already[s] {
			already[s] = true
			nset = append(nset, s)
		}
	}
	for _, solution := range set {
		add("(" + solution + ")")
		for i := range solution {
			add(solution[:i] + "()" + solution[i:])
		}
	}
	caches[n] = nset
	return nset
}
