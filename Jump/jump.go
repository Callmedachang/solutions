package main

import (
	"log"
	"sort"
)

func jump(nums []int) int {
	steps := make([]int, 0)
	return jumpInner(nums, steps, 0, 0)[0]
}
func jumpInner(nums, steps []int, index, step int) []int {
	if index == len(nums)-1 {
		steps = append(steps, step)
	}
	step++
	for i := 0; i < nums[index]; i++ {
		if index+i+1 < len(nums) {
			s := jumpInner(nums, steps, index+i+1, step)
			steps = append(steps, s...)
		}
	}
	sort.Ints(steps)
	return steps
}
func main() {
//[3,4,3,2,5,4,3]
	log.Println(jump([]int{3,4,3,2,5,4,3}))
}
