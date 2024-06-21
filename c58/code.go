package leetcode

func lengthOfLastWord(s string) int {
	length := 0
	n := lengthOfTrimLastSpace(s)
	for i := n - 1; i >= 0; i-- {
		if s[i] == ' ' {
			return length
		}
		length++
	}
	return length
}

func lengthOfTrimLastSpace(s string) int {
	n := len(s)
	for ; n > 0; n-- {
		if s[n-1] != ' ' {
			return n
		}
	}
	return n
}
