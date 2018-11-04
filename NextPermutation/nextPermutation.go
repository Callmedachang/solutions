package main

import (
	"sort"
	"fmt"
)

/*实现获取下一个排列的函数，算法需要将给定数字序列重新排列成字典序中下一个更大的排列。

如果不存在下一个更大的排列，则将数字重新排列成最小的排列（即升序排列）。

必须原地修改，只允许使用额外常数空间。

以下是一些例子，输入位于左侧列，其相应输出位于右侧列。
1,2,3 → 1,3,2
3,2,1 → 1,2,3
1,1,5 → 1,5,1*/
func nextPermutation(nums []int)[]int {
	for i := len(nums) - 1; i > 0; i-- {
		before := nums[i]
		if nums[i-1] >= before {
			continue
		} else {
			sort.Ints(nums[i:])
			for j := i - 1; j >= 0; j-- {
				for k := i; k < len(nums); k++ {
					if nums[j] < nums[k] {
						nums[j], nums[k] = nums[k], nums[j]
						return nums
					}
				}
			}
		}
	}
	sort.Ints(nums)
	return nums
}
func main() {
	t := []int{4,2,4,4,3}
	fmt.Println(nextPermutation(t))
}
