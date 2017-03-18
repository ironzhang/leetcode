package leetcode

func longestPalindrome(s string) string {
	return ""
}

func longestPalindrome_DP(s string) string {
	return ""
}

func longestPalindromeBytes_DFS(b []byte) []byte {
	if isPalindrome(b) {
		return b
	}

	p1 := longestPalindromeBytes_DFS(b[:len(b)-1])
	p2 := longestPalindromeBytes_DFS(b[1:])
	if len(p1) >= len(p2) {
		return p1
	} else {
		return p2
	}
}

func longestPalindromeBytes_BFS(b []byte) []byte {
	var cq, nq [][]byte

	cq = append(cq, b)
	for len(cq) > 0 {
		nq = make([][]byte, 0, len(cq)*2)
		for _, buf := range cq {
			if isPalindrome(buf) {
				return buf
			}

			nq = append(nq, buf[:len(buf)-1])
			nq = append(nq, buf[1:])
		}
		cq = nq
	}

	panic("longestPalindromeBytes_BFS")
}

func isPalindrome(b []byte) bool {
	n := len(b)
	for i := 0; i < n/2; i++ {
		if b[i] != b[n-1-i] {
			return false
		}
	}
	return true
}

func longestPalindrome_Center(s string) string {
	var start, end int
	for i := 0; i < len(s); i++ {
		tmpStart, tmpEnd := maxExpandAroundCenter(s, i)
		if tmpEnd-tmpStart > end-start {
			start, end = tmpStart, tmpEnd
		}
	}
	return s[start : end+1]
}

func maxExpandAroundCenter(s string, i int) (l, r int) {
	l1, r1 := expandAroundCenter(s, i, i)
	l2, r2 := expandAroundCenter(s, i, i+1)
	if r1-l1 >= r2-l2 {
		return l1, r1
	}
	return l2, r2
}

func expandAroundCenter(s string, l, r int) (int, int) {
	for l >= 0 && r < len(s) && s[l] == s[r] {
		l--
		r++
	}
	return l + 1, r - 1
}
