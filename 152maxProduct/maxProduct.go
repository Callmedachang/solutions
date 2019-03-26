package main

import "log"

/*
给定一个整数数组 nums ，找出一个序列中乘积最大的连续子序列（该序列至少包含一个数）。

示例 1:

输入: [2,3,-2,4]
输出: 6
解释: 子数组 [2,3] 有最大乘积 6。
示例 2:

输入: [-2,0,-1]
输出: 0
解释: 结果不能为 2, 因为 [-2,-1] 不是子数组。
*/
func maxProduct(a []int) int {
	cur, neg, max := 1, 1, a[0]

	for i := 0; i < len(a); i++ {

		switch {
		case a[i] > 0:
			cur, neg = a[i]*cur, a[i]*neg
		case a[i] < 0:
			cur, neg = a[i]*neg, a[i]*cur
		default:
			cur, neg = 0, 1
		}

		if max < cur {
			max = cur
		}

		if cur <= 0 {
			cur = 1
		}
	}

	return max
}
func main() {
	//log.Println(maxProduct([]int{-2, -3, 7}))
	//log.Println(maxProduct([]int{2, 3, -2, 4}))
	//log.Println(maxProduct([]int{-2, 0, -1}))
	//log.Println(maxProduct([]int{3, -1, 4}))
	//log.Println(maxProduct([]int{2, -5, -2, -4, 3}))
	log.Println(maxProduct([]int{-1, -2, -3, 0}))
}
