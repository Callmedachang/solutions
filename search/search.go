package main

import "log"

func search(nums []int, target int) int {
	head := 0
	if len(nums) == 1 {
		if nums[head] == target {
			return 0
		} else {
			return -1
		}
	}
	if len(nums) == 0 {
		return -1
	}
	tail := len(nums) - 1
	index := len(nums) / 2
	if target == nums[tail] {
		return tail
	}
	if target == nums[head] {
		return head
	}
	for tail > head+1 {
		if nums[index] == target {
			return index
		}
		if target == nums[tail] {
			return tail
		}
		if target == nums[head] {
			return head
		}
		if target > nums[index] {
			if target <= nums[tail] {
				head = index
				index = (tail-head)/2 + head
				continue
			}
			if target >= nums[tail] && nums[index] < nums[tail] {
				tail = index
				index = (tail-head)/2 + head
				continue
			} else {
				head = index
				index = (tail-head)/2 + head
				continue
			}
		}
		if target < nums[index] {
			if target >= nums[head] {
				tail = index
				index = (tail-head)/2 + head
				continue
			}
			if target <= nums[head] && target <= nums[tail] && nums[index] > nums[head] {
				head = index
				index = (tail-head)/2 + head
				continue
			} else {
				tail = index
				index = (tail-head)/2 + head
				continue
			}
		}
	}
	return -1
}
func main() {
	//[7,8,1,2,3,4,5,6]
	//2
	log.Println(search([]int{2, 3, 4, 5, 6, 7, 8, 9, 1}, 1))
}
