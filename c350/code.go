package leetcode

func intersect(nums1 []int, nums2 []int) []int {
	results := make([]int, 0, len(nums1))

	counter := make(map[int]int, len(nums2))
	for _, n := range nums2 {
		counter[n]++
	}

	for _, n := range nums1 {
		if counter[n] > 0 {
			counter[n]--
			results = append(results, n)
		}
	}

	return results
}
