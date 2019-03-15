package main

import "log"

/*
给定一个未排序的整数数组，找出最长连续序列的长度。

要求算法的时间复杂度为 O(n)。

示例:

输入: [100, 4, 200, 1, 3, 2]
输出: 4
解释: 最长连续序列是 [1, 2, 3, 4]。它的长度为 4。
*/
func longestConsecutive(nums []int) int {
	max := 0
	record := make(map[int]int)
	for _, v := range nums {
		if record[v] != 0 {
			continue
		} else {
			record[v] = 1
			left, right := v, v
			for record[left-1] != 0 {
				left--
			}
			for record[right+1] != 0 {
				right++
			}
			if right-left+1 > max {
				max = right - left + 1
			}
			record[right], record[left] = right-left+1, right-left+1
		}
	}
	return max
}

func main() {
	log.Println(longestConsecutive([]int{1, 3, 5, 2, 4}))
}

/*
some?
利用map的典型案例，利用map起来记录每个连续序列的长度
keyWord: hashMap
*/
