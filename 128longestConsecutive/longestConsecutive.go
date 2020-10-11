package _28longestConsecutive

func LongestConsecutive(nums []int) int {
	res, numSet := 0, make(map[int]*struct{}, 0)
	for _, v := range nums {
		numSet[v] = &struct{}{}
	}
	for k := range numSet {
		if _, ok := numSet[k-1]; !ok {
			currentIndex, leak := k, 1
			if leak > res {
				res = leak
			}
			for numSet[currentIndex+1] != nil {
				delete(numSet,currentIndex+1)
				currentIndex++
				leak++
				if leak > res {
					res = leak
				}
			}
		}
	}
	return res
}
