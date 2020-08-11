package main

import "log"

/*
最长连续序列
给定一个未排序的整数数组，找出最长连续序列的长度。

要求算法的时间复杂度为 O(n)。

示例:

输入: [100, 4, 200, 1, 3, 2]
输出: 4
解释: 最长连续序列是 [1, 2, 3, 4]。它的长度为 4。
*/
func longestConsecutive(nums []int) int {
	nMap := make(map[int]int)
	for _, v := range nums {
		nMap[v] = 1
	}
	res, max := 0, 0
	for k := range nMap {
		index := k
		max = 1
		for nMap[index-1] == 1 {
			index--
			max++
			delete(nMap, index)
		}
		for nMap[k+1] == 1 {
			k++
			max++
			delete(nMap, index)
		}
		if max > res {
			res = max
		}
		max = 1
	}
	return res
}
func main() {
	log.Println(longestConsecutive([]int{100,5, 4, 200, 1, 3, 2}))
}
