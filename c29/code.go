package leetcode

const (
	MaxInt32 = 1<<31 - 1
	MinInt32 = -1 << 31
)

func divide(dividend int, divisor int) int {
	if dividend > MaxInt32 || dividend < MinInt32 {
		return MaxInt32
	}
	if divisor > MaxInt32 || divisor < MinInt32 || divisor == 0 {
		return MaxInt32
	}

	sign := 1
	if dividend < 0 {
		sign = -1 * sign
		dividend = -1 * dividend
	}
	if divisor < 0 {
		sign = -1 * sign
		divisor = -1 * divisor
	}

	division := 0
	for s := divisor; s <= dividend; s += divisor {
		division++
	}
	division = sign * division
	if division > MaxInt32 || division < MinInt32 {
		return MaxInt32
	}
	return division
}
