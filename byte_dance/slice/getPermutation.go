package main

import (
	"math"
	"strconv"
)

/*
给出集合 [1,2,3,…,n]，其所有元素共有 n! 种排列。

按大小顺序列出所有排列情况，并一一标记，当 n = 3 时, 所有排列如下：

"123"
"132"
"213"
"231"
"312"
"321"
给定 n 和 k，返回第 k 个排列。

说明：

给定 n 的范围是 [1, 9]。
给定 k 的范围是[1,  n!]。
示例 1:

输入: n = 3, k = 3
输出: "213"
示例 2:

输入: n = 4, k = 9
输出: "2314"

*/


//1 2 3 4
//2 1 3 4
//2 1 4 3

func getPermutation(n int, k int) string {
	var factors, digits []int // n!的结果 1-1-2-6-24-n! / 数字 [1,2,3,4...n]
	f := 1
	for i := 1; i <= n; i++ {
		factors = append(factors, f)
		digits = append(digits, i)
		f *= i
	}
	var result string
	index, digitIndex := n-1, 0 // 当前的factor索引 / 当前的数字索引
	for index >= 0 {
		// 当余数已经为0时候，后面所有位数直接取digits最右边的数字就好了
		if k == 0 {
			digitIndex = len(digits)
		} else {
			digitIndex = int(math.Ceil(float64(k) / float64(factors[index])))
		}
		result += strconv.Itoa(digits[digitIndex-1])
		k %= factors[index]
		// 取出当前要用的数字
		digits = append(digits[:digitIndex-1], digits[digitIndex:]...)
		index--
	}
	return result
}