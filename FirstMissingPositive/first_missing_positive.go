package main

import (
	"log"
	"sort"
)

func firstMissingPositive(nums []int) int {
	sort.Ints(nums)
	if len(nums) == 0 {
		return 1
	}
	if nums[0] > 1 {
		return 1
	}
	indexOfOne := 0
	hasOne := false
	for i, k := range nums {
		if k == 1 {
			indexOfOne = i
			hasOne = true
			break
		}
	}
	if !hasOne {
		return 1
	}
	for i := indexOfOne; i < len(nums)-1; i++ {
		if nums[i+1]-nums[i] > 1 {
			return nums[i] + 1
		}
	}
	last := nums[len(nums)-1]
	if last < 0 {
		return 1
	}
	return last + 1
}
func main() {
	ss := []int{1000, -1}
	log.Println(firstMissingPositive(ss))
}
