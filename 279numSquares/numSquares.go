package main

import "log"

/*
给定正整数 n，找到若干个完全平方数（比如 1, 4, 9, 16, ...）使得它们的和等于 n。你需要让组成和的完全平方数的个数最少。

示例 1:

输入: n = 12
输出: 3
解释: 12 = 4 + 4 + 4.
示例 2:

输入: n = 13
输出: 2
解释: 13 = 4 + 9.
*/

func numSquares(n int) int {
	res := 0
	for n > 0 {
		n = subMax(n)
		res++
	}
	return res
}
func subMax(n int) int {
	d := 1
	for n >= d*d {
		d++
	}
	return d - (n-1)*(n-1)
}
func main() {
	log.Println(numSquares(12))
}
