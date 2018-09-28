package leetcode

import "testing"

func TestStack(t *testing.T) {
	tests := []rune{'a', 'b', 'c', 'd', '1'}

	s := make(stack, 0, 10)
	for _, r := range tests {
		s.Push(r)
	}
	for i := len(tests) - 1; i >= 0; i-- {
		c, ok := s.Pop()
		if !ok {
			t.Fatalf("stack[%d] pop failed", i)
		}
		if got, want := c, tests[i]; got != want {
			t.Fatalf("stack[%d] pop: got %v, want %v", i, got, want)
		}
	}
	_, ok := s.Pop()
	if got, want := ok, false; got != want {
		t.Fatalf("stack[-1] pop: got %v, want %v", got, want)
	}
}

func TestIsValid(t *testing.T) {
	tests := []struct {
		str   string
		valid bool
	}{
		{"()", true},
		{"()[]{}", true},
		{"(]", false},
		{"([)]", false},
		{"{[]}", true},
	}
	for i, tt := range tests {
		if got, want := isValid(tt.str), tt.valid; got != want {
			t.Errorf("%d: isValid(%q): got %v, want %v", i, tt.str, got, want)
		}
	}
}
