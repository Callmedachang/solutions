package main

import (
	"log"
)

/*
在未排序的数组中找到第 k 个最大的元素。请注意，你需要找的是数组排序后的第 k 个最大的元素，而不是第 k 个不同的元素。

示例 1:

输入: [3,2,1,5,6,4] 和 k = 2
输出: 5
示例 2:

输入: [3,2,3,1,2,4,5,5,6] 和 k = 4
输出: 4
说明:

你可以假设 k 总是有效的，且 1 ≤ k ≤ 数组的长度。
*/

func findKthLargest(nums []int, k int) int {
	if k == 1 {
		max := nums[0]
		for _, v := range nums {
			if v > max {
				max = v
			}
		}
		return max
	}
	res := 0
	big, small := []int{nums[len(nums)-1]}, make([]int, 0)
	equal := true
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] >= nums[len(nums)-1] {
			if nums[i] != big[len(big)-1] {
				equal = false
			}
			big = append(big, nums[i])
		} else {
			small = append(small, nums[i])
		}
	}
	if len(big) >= k {
		if equal {
			return big[0]
		}
		res = findKthLargest(big, k)
	}
	if len(big) < k {
		res = findKthLargest(small, k-len(big))
	}
	return res
}
func main() {
	log.Println(findKthLargest([]int{3,2,1,5,6,4}, 4))
}
