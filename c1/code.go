package leetcode

func twoSum1(nums []int, target int) []int {
	for i, x := range nums {
		for j, y := range nums {
			if i != j && x+y == target {
				return []int{i, j}
			}
		}
	}
	panic("no solution")
}

func twoSum2(nums []int, target int) []int {
	m := make(map[int]int, len(nums))
	for i, x := range nums {
		y := target - x
		if j, ok := m[y]; ok {
			return []int{j, i}
		}
		if _, ok := m[x]; !ok {
			m[x] = i
		}
	}
	panic("no solution")
}
