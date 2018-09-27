package leetcode

import "testing"

func TestCommonPrefix(t *testing.T) {
	tests := []struct {
		a  string
		b  string
		cp string
	}{
		{
			a:  "a",
			b:  "b",
			cp: "",
		},
		{
			a:  "a",
			b:  "a",
			cp: "a",
		},
		{
			a:  "abc",
			b:  "ab",
			cp: "ab",
		},
		{
			a:  "ab",
			b:  "abc",
			cp: "ab",
		},
		{
			a:  "ab123",
			b:  "ab45",
			cp: "ab",
		},
	}
	for i, tt := range tests {
		if got, want := commonPrefix(tt.a, tt.b), tt.cp; got != want {
			t.Errorf("%d: commonPrefix: got %v, want %v", i, got, want)
		}
	}
}

func TestLongestCommonPrefix(t *testing.T) {
	tests := []struct {
		input  []string
		output string
	}{
		{
			input:  []string{"flower", "flow", "flight"},
			output: "fl",
		},
		{
			input:  []string{"dog", "racecar", "car"},
			output: "",
		},
	}
	for i, tt := range tests {
		if got, want := longestCommonPrefix(tt.input), tt.output; got != want {
			t.Errorf("%d: longestCommonPrefix: got %v, want %v", i, got, want)
		}
	}
}
