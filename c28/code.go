package leetcode

func match(s, p string) bool {
	i := len(p)
	if i > len(s) {
		return false
	}
	return s[:i] == p
}

func strStr(haystack string, needle string) int {
	if needle == "" {
		return 0
	}
	for i := 0; i < len(haystack); i++ {
		if match(haystack[i:], needle) {
			return i
		}
	}
	return -1
}
