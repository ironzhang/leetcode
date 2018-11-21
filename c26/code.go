package leetcode

func removeDuplicates(nums []int) int {
	if len(nums) <= 0 {
		return 0
	}

	last := 0
	for i := 1; i < len(nums); i++ {
		if nums[i] != nums[last] {
			last++
			nums[last] = nums[i]
		}
	}
	return last + 1
}
