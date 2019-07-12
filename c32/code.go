package leetcode

type stack []rune

func (s *stack) Push(r rune) {
	*s = append(*s, r)
}

func (s *stack) Pop() (rune, bool) {
	if n := len(*s); n > 0 {
		r := (*s)[n-1]
		*s = (*s)[:n-1]
		return r, true
	}
	return 0, false
}

func isValidParentheses(s string) bool {
	st := make(stack, 0, len(s))
	for _, c := range s {
		switch c {
		case '(':
			st.Push(c)
		case ')':
			r, ok := st.Pop()
			if !ok || r != '(' {
				return false
			}
		default:
			panic("invalid string")
		}
	}
	return len(st) == 0
}

func longestValidParentheses(s string) int {
	return longestValidParenthesesExhaustion(s)
}

func longestValidParenthesesDP(s string) int {
	return 0
}

func longestValidParenthesesExhaustion(s string) int {
	n := len(s)
	max := 0
	for i := 0; i < n; i++ {
		for j := i + 2; j <= n; j += 2 {
			if isValidParentheses(s[i:j]) {
				if x := j - i; x > max {
					max = x
				}
			}
		}
	}
	return max
}

func longestValidParenthesesRecursion(s string) int {
	n := len(s)
	if isValidParentheses(s) {
		return n
	}
	x := longestValidParenthesesRecursion(s[1:])
	y := longestValidParenthesesRecursion(s[:n-1])
	if x > y {
		return x
	} else {
		return y
	}
}
