package main

import (
	"sort"
	"log"
)

func singleNumber(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	if len(nums)==0{
		return 0
	}
	sort.Ints(nums)
	length := len(nums)
	value := 1 << 32
	for i := 0; i < length; i++ {
		if i == length-1 {
			return nums[i]
		}
		if (i+1)%2 == 1 {
			value = nums[i]
			continue
		} else {
			if nums[i] == value {
				value = 1 << 32
				continue
			} else {
				return nums[i-1]
			}
		}
	}
	return 0
}
func main() {
	log.Print(singleNumber([]int{4, 1, 2, 1, 2}))
}
