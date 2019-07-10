package leetcode

func findSubstring(s string, words []string) []int {
	wordCount := len(words)
	if wordCount <= 0 {
		return nil
	}

	wordLength := len(words[0])
	wordHashMap := makeWordHashMap(words)

	var res []int
	for i := 0; i <= len(s)-wordLength*wordCount; i++ {
		subCount := 0
		subHashMap := make(map[string]int)
		for {
			word := s[i+subCount*wordLength : i+(subCount+1)*wordLength]
			subHashMap[word]++
			if subHashMap[word] > wordHashMap[word] {
				break
			}
			subCount++
			if subCount == wordCount {
				res = append(res, i)
				break
			}
		}
	}
	return res
}

func makeWordHashMap(words []string) map[string]int {
	wordHashMap := make(map[string]int)
	for _, word := range words {
		wordHashMap[word]++
	}
	return wordHashMap
}
