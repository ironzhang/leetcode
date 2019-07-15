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
	return longestValidParenthesesStack(s)
}

func longestValidParenthesesDP(s string) int {
	max := 0
	dp := make([]int, len(s))
	for i := 1; i < len(dp); i++ {
		if s[i] == ')' {
			switch s[i-1] {
			case '(':
				if i-2 >= 0 {
					dp[i] = dp[i-2] + 2
				} else {
					dp[i] = 2
				}
			case ')':
				if x := i - dp[i-1] - 1; x >= 0 && s[x] == '(' {
					if y := x - 1; y >= 0 {
						dp[i] = dp[y] + dp[i-1] + 2
					} else {
						dp[i] = dp[i-1] + 2
					}
				}
			}
			if dp[i] > max {
				max = dp[i]
			}
		}
	}
	return max
}

func longestValidParenthesesStack(s string) int {
	max := 0
	st := make([]int, 0, len(s)+1)
	st = append(st, -1)
	for i, c := range s {
		if c == '(' {
			st = append(st, i)
		} else {
			st = st[:len(st)-1]
			if len(st) <= 0 {
				st = append(st, i)
			} else {
				if x := i - st[len(st)-1]; x > max {
					max = x
				}
			}
		}
	}
	return max
}

func longestValidParenthesesLRS(s string) int {
	var left, right, max int
	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			left++
		} else {
			right++
		}
		if left == right {
			if n := right * 2; n > max {
				max = n
			}
		} else if right > left {
			left, right = 0, 0
		}
	}
	left, right = 0, 0
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '(' {
			left++
		} else {
			right++
		}
		if left == right {
			if n := right * 2; n > max {
				max = n
			}
		} else if left > right {
			left, right = 0, 0
		}
	}
	return max
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
