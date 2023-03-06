package leetcode

func merge(nums1 []int, m int, nums2 []int, n int) {
	results := make([]int, 0, m+n)

	i, j := 0, 0
	for {
		if i >= m {
			results = append(results, nums2[j:]...)
			break
		}
		if j >= n {
			results = append(results, nums1[i:]...)
			break
		}

		x, y := nums1[i], nums2[j]
		if x <= y {
			results = append(results, x)
			i++
		} else {
			results = append(results, y)
			j++
		}
	}

	for i, v := range results {
		nums1[i] = v
	}
	return
}
