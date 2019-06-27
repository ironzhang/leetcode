package util

func QuickSort(a []int) {
	if len(a) <= 1 {
		return
	}

	k := partition(a)
	QuickSort(a[:k])
	QuickSort(a[k+1:])
}

func partition(a []int) int {
	t := a[0]
	i, j := 0, len(a)-1
	for i < j {
		for a[j] >= t && i < j {
			j--
		}
		for a[i] <= t && i < j {
			i++
		}
		if i < j {
			a[i], a[j] = a[j], a[i]
		}
	}
	a[0] = a[i]
	a[i] = t
	return j
}
