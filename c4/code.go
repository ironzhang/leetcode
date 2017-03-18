package leetcode

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	sorted := make([]int, 0, len(nums1)+len(nums2))

	i, j := 0, 0
	for {
		if i >= len(nums1) {
			sorted = append(sorted, nums2[j:]...)
			break
		} else if j >= len(nums2) {
			sorted = append(sorted, nums1[i:]...)
			break
		}

		if nums1[i] <= nums2[j] {
			sorted = append(sorted, nums1[i])
			i++
		} else {
			sorted = append(sorted, nums2[j])
			j++
		}
	}

	if len(sorted)%2 == 0 {
		m := len(sorted) / 2
		n := m - 1
		return float64(sorted[m]+sorted[n]) / 2
	} else {
		m := len(sorted) / 2
		return float64(sorted[m])
	}
}

func findMedianSortedArrays2(nums1 []int, nums2 []int) float64 {
}
