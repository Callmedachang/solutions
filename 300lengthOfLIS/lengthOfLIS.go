package _00lengthOfLIS

/*
给定一个无序的整数数组，找到其中最长上升子序列的长度。

示例:

输入: [10,9,2,5,3,7,101,18]
输出: 4
解释: 最长的上升子序列是[2,3,7,101]，它的长度是 4。
说明:

可能会有多种最长上升子序列的组合，你只需要输出对应的长度即可。
你算法的时间复杂度应该为O(n2) 。
进阶: 你能将算法的时间复杂度降低到O(n log n) 吗?
*/

func LengthOfLIS(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	resSli, res := []int{1}, 1
	for i := 1; i < len(nums); i++ {
		temp := 1
		for j := range resSli {
			if nums[i] > nums[j] {
				if resSli[j]+1 > temp {
					temp = resSli[j] + 1
				}
			}
		}
		if temp > res {
			res = temp
		}
		resSli = append(resSli, temp)
	}
	return res
}
