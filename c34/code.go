package leetcode

func searchRange(nums []int, target int) []int {
	return binSearch(nums, 0, len(nums)-1, target)
}

func binSearch(nums []int, left, right, x int) []int {
	for left <= right {
		i := (left + right) / 2
		if nums[i] == x {
			a := leftRangeIndex(nums, left, i, x)
			b := rightRangeIndex(nums, i, right, x)
			return []int{a, b}
		} else if nums[i] > x {
			right = i - 1
		} else {
			left = i + 1
		}
	}
	return []int{-1, -1}
}

func leftRangeIndex(nums []int, left, right, x int) int {
	l, r := left, right
	for l < r {
		if nums[l] == x {
			return left
		}
		i := (l + r) / 2
		if nums[i] < x {
			l = i + 1
		} else {
			r = i
		}
	}
	return right
}

func rightRangeIndex(nums []int, left, right, x int) int {
	for left <= right {
		if nums[right] == x {
			return right
		}
		i := (left + right) / 2
		if nums[i] == x {
			left = i + 1
		} else {
			right = i - 1
		}
	}
	return left
}
