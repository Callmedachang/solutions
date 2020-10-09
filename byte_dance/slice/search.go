package main

func search(nums []int, target int) int {
	if len(nums) == 1 {
		if nums[0] == target {
			return 0
		} else {
			return -1
		}
	}
	left, right := 0, len(nums)-1
	for left < right {
		if left == right-1 {
			if nums[left] == target {
				return left
			}
			if nums[right] == target {
				return right
			}
			return -1
		}
		index := (left + right) / 2
		if nums[index] == target {
			return index
		}
		if nums[left] == target {
			return left
		}
		if nums[right] == target {
			return right
		}
		if index < len(nums)-1 {
			if nums[index+1] == target {
				return index + 1
			}
		}
		if target > nums[index] {
			if target <= nums[right] {
				left = index
				continue
			}
			if target >= nums[right] && nums[index] < nums[right] {
				right = index
				continue
			} else {
				left = index
				continue
			}
		}
		if target < nums[index] {
			if target >= nums[left] {
				right = index
				continue
			}
			if target <= nums[left] && target <= nums[right] && nums[index] > nums[left] {
				left = index
				continue
			} else {
				right = index
				continue
			}
		}
	}
	return -1
}
