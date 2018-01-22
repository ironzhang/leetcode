package leetcode

func myAtoi(str string) int {
	sign, i := getSign(str)
	if sign > 0 {
		return positiveAtoi(str[i:])
	} else if sign < 0 {
		return negativeAtoi(str[i:])
	} else {
		return 0
	}
}

func getSign(str string) (sign, i int) {
	for i, c := range str {
		switch c {
		case ' ':
			continue
		case '+':
			return 1, i + 1
		case '-':
			return -1, i + 1
		default:
			return 1, i
		}
	}
	return 0, 0
}

func positiveAtoi(str string) int {
	num := 0
	for _, c := range str {
		if !(c >= '0' && c <= '9') {
			return num
		}
		num = num*10 + int(c-'0')
		if num > 2147483647 {
			return 2147483647
		}
	}
	return num
}

func negativeAtoi(str string) int {
	num := 0
	for _, c := range str {
		if !(c >= '0' && c <= '9') {
			return -num
		}
		num = num*10 + int(c-'0')
		if num > 2147483648 {
			return -2147483648
		}
	}
	return -num
}
