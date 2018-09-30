package main

import (
	"sort"
	"fmt"
)

func main() {
	nums := []int{1, 2, 5, 13, -1, 3}
	fmt.Println(threeSumClosest(nums, 5))
}

func threeSumClosest(nums []int, target int) int {
	sort.Ints(nums)
	arrLen := len(nums)
	result := nums[0] + nums[1] + nums[arrLen-1]
	var sum int //三个数组相加的值
	if len(nums) < 3 {
		return 0
	}

	for i := 0; i < arrLen; i++ {
		low, high := i+1, arrLen-1
		for low < high {
			sum = nums[i] + nums[low] + nums[high]
			if sum > target {
				high--
			} else {
				low++
			}
			if abs(sum-target) < abs(result-target) {
				result = sum
			}
		}
	}

	return result

}

//求绝对值
func abs(a int) int {
	if a < 0 {
		return -a
	}
	if a == 0 {
		return 0
	}
	return a
}