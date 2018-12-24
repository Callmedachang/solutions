package main

import "log"

/*
实现 pow(x, n) ，即计算 x 的 n 次幂函数。

示例 1:

输入: 2.00000, 10
输出: 1024.00000
示例 2:

输入: 2.10000, 3
输出: 9.26100
示例 3:

输入: 2.00000, -2
输出: 0.25000
解释: 2-2 = 1/22 = 1/4 = 0.25
说明:

-100.0 < x < 100.0
n 是 32 位有符号整数，其数值范围是 [−231, 231 − 1] 。
*/
//TODO:若果当前算法复杂度是O(n)那么就要考虑是否可以向log(n)靠近
func myPow(x float64, n int) float64 {
	res := x
	if n == 0 || x == 1 {
		return 1
	}
	if x == 0 {
		return 0
	}
	needRe := false
	if n < 0 {
		needRe = true
		n = 0 - n
	}
	res = pow(x, n)
	if needRe {
		res = 1 / res
	}
	return res
}
func pow(x float64, n int) float64 {
	if n == 1 {
		return x
	}
	if n == 0 {
		return 1
	}
	res := pow(x, n>>1)
	if n%2 == 1 {
		return res * res * x
	}
	return res * res
}

func main() {
	log.Println(myPow(2, 2))
}
