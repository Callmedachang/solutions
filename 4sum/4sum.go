package main

import (
	"fmt"
	"sort"
)

/*
给定一个包含 n 个整数的数组 nums 和一个目标值 target，判断 nums 中是否存在四个元素 a，b，c 和 d ，使得 a + b + c + d 的值与 target 相等？找出所有满足条件且不重复的四元组。

注意：

答案中不可以包含重复的四元组。

示例：

给定数组 nums = [1, 0, -1, 0, -2, 2]，和 target = 0。

满足要求的四元组集合为：
[
  [-1,  0, 0, 1],
  [-2, -1, 1, 2],
  [-2,  0, 0, 2]
]*/
func fourSum(nums []int, target int) [][]int {
	resOut := make([][]int, 0)
	for i := 0; i < len(nums)-3; i++ {
		first := nums[i]
		res := threeSum(nums[i+1:], target-first)
		for _, v := range res {
			v := append(v, nums[i])
			sort.Ints(v)
			has := false
			for _, inResOut := range resOut {
				if inResOut[0] == v[0] && inResOut[1] == v[1] && inResOut[2] == v[2] {
					has = true
				}
			}
			if !has {
				resOut = append(resOut, v)
			}
		}
	}
	return resOut
}
func threeSum(nums []int, targetIn int) [][]int {
	var ret [][]int
	if len(nums) < 3 {
		return ret
	}
	sort.Ints(nums)
	for i := 0; i < len(nums); i++ {
		if i != 0 && nums[i] == nums[i-1] {
			continue
		}
		target := targetIn - nums[i]
		for j := i + 1; j < len(nums); j++ {
			if j > i+1 && nums[j] == nums[j-1] {
				continue
			}
			hit := target - nums[j]
			if hit < nums[j] {
				continue
			}
			for k := j + 1; k < len(nums) && hit >= nums[k]; k++ {
				if hit == nums[k] {
					tmp := []int{nums[i], nums[j], nums[k]}
					ret = append(ret, tmp)
					break
				}
			}
		}
	}
	return ret
}
func main() {
	//nums := []int{1, 0, -1, 0, -2, 2}
	//nums := []int{0, 0, 0, 0}
	nums := []int{5,5,3,5,1,-5,1,-2}
	fmt.Println(fourSum(nums, 0))
}