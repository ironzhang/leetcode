package leetcode

import (
	"testing"
)

func TestLongestValidParentheses(t *testing.T) {
	tests := []struct {
		input  string
		output int
	}{
		{
			input:  "()(()",
			output: 2,
		},
		{
			input:  ")()())",
			output: 4,
		},
		{
			input:  "(()",
			output: 2,
		},
		{
			input:  ")(()(()(((())(((((()()))((((()()(()()())())())()))()()()())(())()()(((()))))()((()))(((())()((()()())((())))(())))())((()())()()((()((())))))((()(((((()((()))(()()(())))((()))()))())",
			output: 132,
		},
	}
	for i, tt := range tests {
		n := longestValidParentheses(tt.input)
		if got, want := n, tt.output; got != want {
			t.Errorf("%d: got %v, want %v", i, got, want)
		} else {
			t.Logf("%d: got %v", i, got)
		}
	}
}
