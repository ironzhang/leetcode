package leetcode

import "fmt"

func getFolderNames(names []string) []string {
	counter := make(map[string]int)
	for i, name := range names {
		k := counter[name]
		if k <= 0 {
			counter[name] = k + 1
			continue
		}

		for {
			kname := fmt.Sprintf("%s(%d)", name, k)
			k++
			if _, ok := counter[kname]; !ok {
				names[i] = kname
				counter[kname] = 1
				break
			}
		}
		counter[name] = k
	}
	return names
}
