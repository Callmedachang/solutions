package _1nextPermutation

import (
	"sort"
)

/*
实现获取下一个排列的函数，算法需要将给定数字序列重新排列成字典序中下一个更大的排列。

如果不存在下一个更大的排列，则将数字重新排列成最小的排列（即升序排列）。

必须原地修改，只允许使用额外常数空间。

以下是一些例子，输入位于左侧列，其相应输出位于右侧列。
1,2,3 → 1,3,2
3,2,1 → 1,2,3
1,1,5 → 1,5,1
*/

func NextPermutation(nums []int) {
	numsMap := make(map[int]int)
	max := 0
	for i := len(nums) - 1; i > 0; i-- {
		if nums[i] > max {
			max = nums[i]
		}
		numsMap[nums[i]] = i
		index, has := findLess(nums[i-1], numsMap,max)
		if has {
			nums[i-1], nums[index] = nums[index], nums[i-1]
			sort.Ints(nums[i:])
			return
		}
	}
	head, tail := 0, len(nums)-1
	for head < tail {
		nums[head], nums[tail] = nums[tail], nums[head]
		head++
		tail--
	}
	return
}

func findLess(num int, m map[int]int, max int) (int, bool) {
	for num < max {
		index, has := m[num+1]
		if has {

			return index, has
		}
		num++
	}
	return 0, false
}
