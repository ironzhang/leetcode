package leetcode

import (
	"testing"
)

func TestMatch(t *testing.T) {
	tests := []struct {
		s string
		p string
		b bool
	}{
		{
			s: "",
			p: "",
			b: true,
		},
		{
			s: "a",
			p: "",
			b: true,
		},
		{
			s: "a",
			p: "a",
			b: true,
		},
		{
			s: "",
			p: "a",
			b: false,
		},
		{
			s: "aa",
			p: "a",
			b: true,
		},
		{
			s: "aa",
			p: "aa",
			b: true,
		},
		{
			s: "aa",
			p: "aaa",
			b: false,
		},
		{
			s: "aa",
			p: "bb",
			b: false,
		},
	}
	for i, tt := range tests {
		b := match(tt.s, tt.p)
		if got, want := b, tt.b; got != want {
			t.Fatalf("%d: got %v, want %v", i, got, want)
		} else {
			t.Logf("%d: got %v", i, got)
		}
	}
}

func TestStrStr(t *testing.T) {
	tests := []struct {
		haystack string
		needle   string
		index    int
	}{
		{
			haystack: "hello",
			needle:   "ll",
			index:    2,
		},
		{
			haystack: "aaaaa",
			needle:   "bba",
			index:    -1,
		},
		{
			haystack: "aaaaa",
			needle:   "",
			index:    0,
		},
		{
			haystack: "",
			needle:   "aaaaa",
			index:    -1,
		},
		{
			haystack: "",
			needle:   "",
			index:    0,
		},
	}
	for i, tt := range tests {
		n := strStr(tt.haystack, tt.needle)
		if got, want := n, tt.index; got != want {
			t.Fatalf("%d: got %v, want %v", i, got, want)
		} else {
			t.Logf("%d: got %v", i, got)
		}
	}
}
