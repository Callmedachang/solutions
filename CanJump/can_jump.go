package main

import "log"

func canJump(nums []int) bool {
	return jump(nums, len(nums)-1)
}
func jump(nums []int, index int) bool {
	res := false
	if index == 0 {
		res = true
		return true
	}
	for i := index - 1; i >= 0; i-- {
		if nums[i] >= (index - i) {
			if jump(nums, i) {
				return true
			}
		}
	}
	return res
}
func main() {
	s := []int{2,0,6,9,8,4,5,0,8,9,1,2,9,6,8,8,0,6,3,1,2,2,1,2,6,5,3,1,2,2,6,4,2,4,3,0,0,0,3,8,2,4,0,1,2,0,1,4,6,5,8,0,7,9,3,4,6,6,5,8,9,3,4,3,7,0,4,9,0,9,8,4,3,0,7,7,1,9,1,9,4,9,0,1,9,5,7,7,1,5,8,2,8,2,6,8,2,2,7,5,1,7,9,6}
	log.Println(canJump(s))
}
