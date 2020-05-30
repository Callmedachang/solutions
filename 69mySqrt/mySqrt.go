package main

import "log"

/*
实现 int sqrt(int x) 函数。

计算并返回 x 的平方根，其中 x 是非负整数。

由于返回类型是整数，结果只保留整数的部分，小数部分将被舍去。

示例 1:

输入: 4
输出: 2
示例 2:

输入: 8
输出: 2
说明: 8 的平方根是 2.82842...,
     由于返回类型是整数，小数部分将被舍去。

*/

func mySqrt(x int) int {
	head, tail := 0, x
	res := x / 2
	for !isSqrt(res, x) {
		if res*res < x {
			if (res+1)*(res+1) == x {
				return res + 1
			}
			head = res
			res = (res + tail) / 2
		} else {
			if (res-1)*(res-1) == x {
				return res - 1
			}
			tail = res
			res = (res + head) / 2
		}
	}
	return res
}
func isSqrt(s, m int) bool {
	if (s*s <= m) && ((s+1)*(s+1) > m) {
		return true
	} else {
		return false
	}
}
func main() {
	log.Println(mySqrt(121))
}
