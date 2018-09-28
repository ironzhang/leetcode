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

var brackets = map[rune]rune{
	'(': ')',
	'[': ']',
	'{': '}',
}

func isValid(s string) bool {
	var st = make(stack, 0, len(s))
	for _, c := range s {
		switch c {
		case '(', '[', '{':
			st.Push(c)
		case ')', ']', '}':
			r, ok := st.Pop()
			if !ok || brackets[r] != c {
				return false
			}
		default:
			panic("invalid params")
		}
	}
	return len(st) == 0
}
