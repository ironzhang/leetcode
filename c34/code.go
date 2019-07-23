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

func leftRangeIndex(nums []int, l, r, x int) int {
	for l <= r {
		if nums[l] == x {
			return l
		}
		i := (l + r) / 2
		if nums[i] < x {
			l = i + 1
		} else if nums[i] == x {
			r = i
		} else {
			return -1
		}
	}
	return -1
}

func rightRangeIndex(nums []int, l, r, x int) int {
	for l <= r {
		if nums[r] == x {
			return r
		}
		i := (l + r + 1) / 2
		if nums[i] == x {
			l = i
		} else if nums[i] > x {
			r = i - 1
		} else {
			return -1
		}
	}
	return -1
}
