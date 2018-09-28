package leetcode

var m = map[rune]string{
	'2': "abc",
	'3': "def",
	'4': "ghi",
	'5': "jkl",
	'6': "mon",
	'7': "pqrs",
	'8': "tuv",
	'9': "wxyz",
}

func letterCombinations(digits string) []string {
	if digits == "" {
		return []string{}
	}
	prev := []string{""}
	for _, d := range digits {
		next := make([]string, 0, len(prev)*len(m[d]))
		for _, s := range prev {
			for _, c := range m[d] {
				next = append(next, s+string(c))
			}
		}
		prev = next
	}
	return prev
}
