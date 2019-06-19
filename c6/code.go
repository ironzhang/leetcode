package leetcode

func convert(s string, numRows int) string {
	if numRows == 1 {
		return s
	}

	m := make([][]rune, numRows)

	i := 0
	forward := true
	for _, r := range s {
		m[i] = append(m[i], r)
		if forward {
			i++
			if i == numRows {
				i -= 2
				forward = false
			}
		} else {
			i--
			if i == -1 {
				i += 2
				forward = true
			}
		}
	}

	var ss string
	for _, row := range m {
		ss += string(row)
	}
	return ss
}
