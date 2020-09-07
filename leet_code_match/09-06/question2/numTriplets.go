package question2

import (
	"sort"
)

func NumTriplets(nums1 []int, nums2 []int) int {
	sort.Ints(nums1)
	sort.Ints(nums2)
	res, triNums1, triNums2 := 0, make([]int, len(nums1)), make([]int, len(nums2))
	for i, v := range nums1 {
		triNums1[i] = v * v
	}
	for i, v := range nums2 {
		triNums2[i] = v * v
	}
	res += findTwo(triNums1, nums2)
	res += findTwo(triNums2, nums1)
	return res
}

func findTwo(triNums1, nums2 []int) int {
	res := 0
	for i := 0; i < len(triNums1); i++ {
		for j := 0; j < len(nums2); j++ {
			head, tail := j, len(nums2)-1
			min, max := head+1, len(nums2)-1
			for   head < tail {
				if triNums1[i] == nums2[head]*nums2[tail] {
					res++
					tempTail1, tempTail2 := tail-1, tail+1
					for head < tempTail1 {
						if triNums1[i] == nums2[head]*nums2[tempTail1] {
							res++
							tempTail1--
						} else {
							break
						}
					}
					for tempTail2 < len(triNums1)-1 {
						if triNums1[i] == nums2[head]*nums2[tempTail2] {
							res++
							tempTail2++
						} else {
							break
						}
					}
					break
				}
				if triNums1[i] > nums2[head]*nums2[tail] {
					min = tail
					if tail == (min + max) / 2{
						break
					}
					tail = (min + max) / 2
				} else {
					max = tail
					if tail == (min + max) / 2{
						break
					}
					tail = (min + max) / 2
				}
			}
		}
	}
	return res
}
