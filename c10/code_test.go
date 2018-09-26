package leetcode

import (
	"testing"
)

func TestIsMatch(t *testing.T) {
	tests := []struct {
		s string
		p string
		r bool
	}{
		{"a", "a", true},
		{"aa", "a", false},
		{"aa", "aa", true},
		{"aaa", "aa", false},
		{"aa", "a*", true},
		{"aa", ".*", true},
		{"ab", ".*", true},
		{"aab", "c*a*b", true},
		{"mississippi", "mis*is*ip*.", true},
		{"aaa", "a*a", true},
		{"aaba", "ab*a*c*a", false},
		{"", "c*c*", true},
		{"a", "", false},
		{"", "", true},
	}
	for i, tt := range tests {
		if got, want := isMatch(tt.s, tt.p), tt.r; got != want {
			m, _ := re2machine(tt.p)
			t.Errorf("%d: isMatch(%q, %q): got=%v, want=%v, machine:\n%s", i, tt.s, tt.p, got, want, m.String())
		}
	}
}
