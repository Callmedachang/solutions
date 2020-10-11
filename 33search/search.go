package _3search

/*
假设按照升序排序的数组在预先未知的某个点上进行了旋转。

( 例如，数组 [0,1,2,4,5,6,7] 可能变为 [4,5,6,7,0,1,2] )。

搜索一个给定的目标值，如果数组中存在这个目标值，则返回它的索引，否则返回 -1 。

你可以假设数组中不存在重复的元素。

你的算法时间复杂度必须是 O(log n) 级别。

示例 1:

输入: nums = [4,5,6,7,0,1,2], target = 0
输出: 4
示例 2:

输入: nums = [4,5,6,7,0,1,2], target = 3
输出: -1
*/

func Search(nums []int, target int) int {
	if len(nums) == 1 {
		if nums[0] == target {
			return 0
		} else {
			return -1
		}
	}
	head, tail := 0, len(nums)-1
	for head <= tail {
		mid := (head + tail) / 2
		if nums[mid] == target {
			return mid
		}
		if nums[0] <= nums[mid] { //如果在前半个序列
			if nums[0] <= target && target < nums[mid] {
				tail = mid - 1
			} else {
				head = mid + 1
			}
		} else { //如果在后半个序列
			if nums[mid] < target && target <= nums[len(nums)-1] {
				head = mid + 1
			} else {
				tail = mid - 1
			}
		}
	}
	return -1
}
