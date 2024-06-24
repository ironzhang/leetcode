package leetcode

import "testing"

func TestEncode(t *testing.T) {
	tests := []struct {
		input  string
		output string
	}{
		{
			input:  "1",
			output: "11",
		},
		{
			input:  "11",
			output: "21",
		},
		{
			input:  "21",
			output: "1211",
		},
	}
	for i, tt := range tests {
		output := encode(tt.input)
		if got, want := output, tt.output; got != want {
			t.Errorf("testcase(%d): encode: %v != %v", i, got, want)
		}
	}
}

func TestCountAndSay(t *testing.T) {
	tests := []struct {
		n int
		s string
	}{
		{
			n: 1,
			s: "1",
		},
		{
			n: 2,
			s: "11",
		},
		{
			n: 3,
			s: "21",
		},
		{
			n: 4,
			s: "1211",
		},
	}
	for i, tt := range tests {
		s := countAndSay(tt.n)
		if got, want := s, tt.s; got != want {
			t.Errorf("testcase(%d): countAndSay: %v != %v", i, got, want)
		}
	}
}
