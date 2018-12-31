package main

import "log"

func canJump2(nums []int) bool {
	if len(nums) < 2 {
		return true
	}

	for i := 0; i < len(nums)-1; i++ {
		if nums[i] == 0 {
			if !canJumpZero(nums, i) {
				return false
			}
		}
	}
	return true
}
func canJumpZero(nums []int, index int) bool {
	for i := index - 1; i >= 0; i-- {
		if nums[i] > index-i {
			return true
		}
	}
	return false
}
func main() {
	//s := []int{3,2,1,0,4}
	s := []int{2, 0, 0}
	log.Println(canJump2(s))
}
