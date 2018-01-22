package leetcode

import "fmt"

func convert(s string, numRows int) string {
	n := 0
	b := make([]rune, len(s))
	for _, c := range s {
		fmt.Printf("%c\n", c)
	}
	return s
}
