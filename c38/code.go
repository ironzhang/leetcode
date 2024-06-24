package leetcode

import (
	"fmt"
	"strings"
)

func encode(s string) string {
	var b strings.Builder

	dup := 0
	last := len(s) - 1
	for i := 0; i <= last; i++ {
		dup++
		if i == last || s[i] != s[i+1] {
			fmt.Fprintf(&b, "%d%c", dup, s[i])
			dup = 0
		}
	}
	return b.String()
}

func countAndSay(n int) string {
	s := "1"
	for i := 2; i <= n; i++ {
		s = encode(s)
	}
	return s
}
