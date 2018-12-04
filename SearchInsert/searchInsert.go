package main

import "fmt"

func searchInsert(nums []int, target int) int {
	head := 0
	tail := len(nums) - 1
	mid := 0
	for head <= tail {
		mid = (head + tail) / 2
		if nums[mid] == target {
			return mid
		} else if nums[mid] > target {
			tail = mid
		} else if nums[mid] < target {
			head = mid
		}
		if tail-head < 2 {
			if target <= nums[head] {
				mid = head
			}
			if target >= nums[tail] {
				mid = tail
			}
			break
		}
	}
	if nums[mid] >= target {
		return mid
	} else {
		return mid + 1
	}
}

func main() {
	s2 := []int{1,3,5}
	fmt.Println(searchInsert(s2, 1))
}
