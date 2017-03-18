package leetcode

import "testing"

func TestLengthOfLongestSubstring(t *testing.T) {
	tests := []struct {
		str  string
		want int
	}{
		{
			str:  "abcabcbb",
			want: 3,
		},
		{
			str:  "bbbbb",
			want: 1,
		},
		{
			str:  "pwwkew",
			want: 3,
		},
	}

	for i, test := range tests {
		if got, want := lengthOfLongestSubstring(test.str), test.want; got != want {
			t.Errorf("testcase(%d): %v != %v", i, got, want)
		}
	}

}

func TestLongestSubstring(t *testing.T) {
	tests := []struct {
		input  string
		output string
	}{
		{
			input:  "abcabcbb",
			output: "abc",
		},
		{
			input:  "bbbbb",
			output: "b",
		},
		{
			input:  "pwwkew",
			output: "wke",
		},
		{
			input:  "pwwk",
			output: "pw",
		},
	}

	for i, test := range tests {
		if got, want := longestSubstring([]byte(test.input)), test.output; string(got) != want {
			t.Errorf("testcase(%d): %v != %v", i, string(got), want)
		}
	}
}
