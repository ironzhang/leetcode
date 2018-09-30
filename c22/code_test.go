package leetcode

import (
	"reflect"
	"testing"
)

func slice2map(s []string) map[string]bool {
	m := make(map[string]bool)
	for _, k := range s {
		m[k] = true
	}
	return m
}

func deepEqual(s1, s2 []string) bool {
	m1 := slice2map(s1)
	m2 := slice2map(s2)
	return reflect.DeepEqual(m1, m2)
}

func TestGenerateParenthesis(t *testing.T) {
	tests := []struct {
		n int
		s []string
	}{
		{
			n: 0,
			s: []string{},
		},
		{
			n: 1,
			s: []string{"()"},
		},
		{
			n: 2,
			s: []string{"(())", "()()"},
		},
		{
			n: 3,
			s: []string{
				"((()))",
				"()(())",
				"(())()",
				"(()())",
				"()()()",
			},
		},
		{
			n: 4,
			s: []string{
				"(((())))",
				"()((()))",
				"((()))()",
				"(()(()))",
				"()()(())",
				"()(())()",
				"((())())",
				"(())()()",
				"((()()))",
				"()(()())",
				"(()())()",
				"(()()())",
				"()()()()",
				"(())(())",
			},
		},
	}
	for i, tt := range tests {
		if got, want := generateParenthesis(tt.n), tt.s; !deepEqual(got, want) {
			t.Errorf("%d: generateParenthesis(%d): got %v, want %v", i, tt.n, got, want)
		}
	}
}
