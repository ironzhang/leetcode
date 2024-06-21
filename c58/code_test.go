package leetcode

import "testing"

func TestLengthOfTrimLastSpace(t *testing.T) {
	tests := []struct {
		s string
		n int
	}{
		{
			s: "",
			n: 0,
		},
		{
			s: " a",
			n: 2,
		},
		{
			s: " a ",
			n: 2,
		},
		{
			s: " a b",
			n: 4,
		},
		{
			s: " a b  ",
			n: 4,
		},
	}
	for i, tt := range tests {
		n := lengthOfTrimLastSpace(tt.s)
		if got, want := n, tt.n; got != want {
			t.Errorf("testcase(%d): lengthOfTrimLastSpace: %v != %v", i, got, want)
		}
	}
}

func TestLengthOfLastWord(t *testing.T) {
	tests := []struct {
		s string
		n int
	}{
		{
			s: "",
			n: 0,
		},
		{
			s: "Hello World",
			n: 5,
		},
		{
			s: "   fly me   to   the moon  ",
			n: 4,
		},
		{
			s: "luffy is still joyboy",
			n: 6,
		},
	}
	for i, tt := range tests {
		n := lengthOfLastWord(tt.s)
		if got, want := n, tt.n; got != want {
			t.Errorf("testcase(%d): lengthOfLastWord: %v != %v", i, got, want)
		}
	}
}
