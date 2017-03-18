package leetcode

func lengthOfLongestSubstring(s string) int {
	i := 0
	start := 0
	sublen := 0
	longest := 0
	characters := make(map[byte]int)

	for i < len(s) {
		c := s[i]
		if pos, ok := characters[c]; ok && pos >= start {
			sublen = i - start
			if sublen > longest {
				longest = sublen
			}
			start++
			continue
		}
		characters[c] = i
		i++
	}

	sublen = i - start
	if sublen > longest {
		longest = sublen
	}

	return longest
}

func longestSubstring(s []byte) []byte {
	start := 0
	sublen := 0
	longest := 0
	characters := make(map[byte]int)
	var substr []byte

	for i := 0; i < len(s); i++ {
		c := s[i]
		if pos, ok := characters[c]; ok && pos >= start {
			sublen = i - start
			if sublen > longest {
				longest = sublen
				substr = s[start:i]
			}
			start++
			i--
			continue
		}
		characters[c] = i
	}

	sublen = len(s) - start
	if sublen > longest {
		longest = sublen
		substr = s[start:]
	}

	return substr
}
