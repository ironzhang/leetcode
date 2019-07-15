package leetcode

func search(nums []int, target int) int {
	if len(nums) <= 0 {
		return -1
	}

	index := searchRotatedIndex(nums)
	if target >= nums[0] {
		return binSearch(nums, 0, index-1, target)
	} else {
		return binSearch(nums, index, len(nums)-1, target)
	}
}

func searchRotatedIndex(nums []int) int {
	n := len(nums)
	x := nums[0]
	i, left, right := 0, 1, n-1
	for left <= right {
		i = (left + right) / 2
		if nums[i] < nums[i-1] {
			return i
		}
		if nums[i] > x {
			left = i + 1
		} else {
			right = i - 1
		}
	}
	return n
}

func binSearch(nums []int, left, right, x int) int {
	for left <= right {
		i := (left + right) / 2
		if nums[i] == x {
			return i
		} else if nums[i] > x {
			right = i - 1
		} else {
			left = i + 1
		}
	}
	return -1
}
