package main

func findLengthOfLCIS(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	temp, max := 1, 1
	for i := 1; i < len(nums); i++ {
		if nums[i] > nums[i-1] {
			temp++
		} else {
			if temp > max {
				max = temp
			}
			temp = 1
		}
	}
	if temp > max {
		max = temp
	}
	return max
}
