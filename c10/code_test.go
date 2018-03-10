package leetcode

import "testing"

func TestIsMatch(t *testing.T) {
	tests := []struct {
		s string
		p string
		r bool
	}{
		{"aa", "a", false},
		{"aa", "aa", true},
		{"aaa", "aa", false},
		{"aa", "a*", true},
		{"aa", ".*", true},
		{"ab", ".*", true},
		{"aab", "c*a*b", true},
	}
	for _, tt := range tests {
		if got, want := isMatch(tt.s, tt.p), tt.r; got != want {
			t.Errorf("isMatch(%q, %q): got=%v, want=%v", tt.s, tt.p, got, want)
		}
	}
}
