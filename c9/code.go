package leetcode

func isPalindrome2(x int) bool {
	if x < 0 {
		return false
	}

	y, z := 0, x
	for x > 0 {
		y = y*10 + x%10
		x = x / 10
	}
	return y == z
}

func isPalindrome(x int) bool {
	if x < 0 || (x%10 == 0 && x != 0) {
		return false
	}

	y := 0
	for x > y {
		y = y*10 + x%10
		x = x / 10
	}
	return (x == y) || (x == y/10)
}
