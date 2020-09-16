package main

import "sort"

//[-2,-1,0,1,2]

func threeSum2(nums []int,target int) [][]int {
	// 排序后，可以按规律查找
	sort.Ints(nums)
	res := make([][]int, 0)
	for i := range nums {
		// 避免添加重复的结果
		// i>0 是为了防止nums[i-1]溢出
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		left, right := i+1, len(nums)-1

		for left < right {
			sum := nums[i] + nums[left] + nums[right]
			switch {
			case sum < target:
				// 较小的 left 需要变大
				left++
			case sum > target:
				// 较大的 right 需要变小
				right--
			default:
				res = append(res, []int{nums[i], nums[left], nums[right]})
				// 为避免重复添加，left 和 right 都需要移动到不同的元素上。
				left, right = next(nums, left, right)
			}
		}
	}
	return res
}

func next(nums []int, l, r int) (int, int) {
	for l < r {
		switch {
		case nums[l] == nums[l+1]:
			l++
		case nums[r] == nums[r-1]:
			r--
		default:
			l++
			r--
			return l, r
		}
	}
	return l, r
}
