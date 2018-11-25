package main

import "fmt"

/*给定一个按照升序排列的整数数组 nums，和一个目标值 target。找出给定目标值在数组中的开始位置和结束位置。

你的算法时间复杂度必须是 O(log n) 级别。

如果数组中不存在目标值，返回 [-1, -1]。

示例 1:

输入: nums = [5,7,7,8,8,10], target = 8
输出: [3,4]
示例 2:

输入: nums = [5,7,7,8,8,10], target = 6
输出: [-1,-1]*/
func searchRange(nums []int, target int) []int {
	if len(nums)==0{
		return []int{-1,-1}
	}
	before := 0
	end := 0
	length := len(nums)
	head := 0
	tail := length - 1
	for tail > head {
		mid := (head + tail) / 2
		if nums[mid] == target {
			before = mid
			end = mid
			break
		} else if nums[mid] > target {
			tail = mid
		} else {
			head = mid
		}
	}

	for before > head {
		index := before-1
		if nums[index] == target {
			before--
		}else {
			break
		}
	}
	for end < tail {
		index := end+1
		if nums[index] == target {
			end++
		}else {
			break
		}
	}
	return []int{before, end}
}
func main() {
	//ss := []int{5, 7, 7, 8, 8, 10}
	//fmt.Println(searchRange(ss, 8))
	s2 := []int{10}
	fmt.Println(searchRange(s2, 10))
}
