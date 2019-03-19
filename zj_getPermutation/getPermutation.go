package main

import (
	"log"
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
输出: "2314"*/

func getPermutation(n int, k int) string {
	res := ""
	has := make(map[int]int, n)

	for n > 0 {
		em := getEveryMul(n)
		fir := 1
		if k > em {
			fir = k / em
			k = k % em
			fir++
		} else {
			fir = 1
		}
		if fir > 1 && k == 0 {
			fir--
			k = em
		}
		for has[fir] != 0 {
			fir++
		}
		res = res + strconv.Itoa(fir)
		has [fir] = 1
		n --
	}
	return res
}

func getEveryMul(n int) int {
	res := 1
	for i := 1; i < n; i++ {
		res = res * i
	}
	return res
}

func main() {
	log.Println(getPermutation(3, 2))
}
