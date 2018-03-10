package leetcode

import "testing"

func TestIsPalindrome(t *testing.T) {
	tests := []struct {
		x int
		r bool
	}{
		{x: 1, r: true},
		{x: 11, r: true},
		{x: 101, r: true},
		{x: 10, r: false},
	}
	for _, tt := range tests {
		if got, want := isPalindrome(tt.x), tt.r; got != want {
			t.Errorf("is palindrome: x=%d, got=%t, want=%t", tt.x, got, want)
		}
	}
}
