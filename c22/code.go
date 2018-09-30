package leetcode

func generateParenthesis(n int) []string {
	return backtrack(nil, "", 0, 0, n)
}

func backtrack(ans []string, s string, open, close, max int) []string {
	if len(s) == max*2 {
		ans = append(ans, s)
		return ans
	}
	if open < max {
		ans = backtrack(ans, s+"(", open+1, close, max)
	}
	if close < open {
		ans = backtrack(ans, s+")", open, close+1, max)
	}
	return ans
}

func generateParenthesis0(n int) []string {
	if n == 0 {
		return []string{}
	} else if n == 1 {
		return []string{"()"}
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
	return nset
}
