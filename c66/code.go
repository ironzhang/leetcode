package leetcode

func plusOne(digits []int) []int {
	n := len(digits)
	borrow := 1
	results := make([]int, n+1)
	for i := n - 1; i >= 0; i-- {
		s := borrow + digits[i]
		borrow = s / 10
		results[i+1] = s % 10
	}
	if borrow > 0 {
		results[0] = borrow
		return results
	}
	return results[1:]
}
