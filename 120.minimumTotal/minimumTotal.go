package main

import "log"

/*
给定一个三角形，找出自顶向下的最小路径和。每一步只能移动到下一行中相邻的结点上。

例如，给定三角形：

[
     [2],
    [3,4],
   [6,5,7],
  [4,1,8,3]
]
自顶向下的最小路径和为 11（即，2 + 3 + 5 + 1 = 11）。

说明：

如果你可以只使用 O(n) 的额外空间（n 为三角形的总行数）来解决这个问题，那么你的算法会很加分。
*/
func minimumTotal(triangle [][]int) int {
	for i := 0; i < len(triangle); i++ {
		for j := 0; j < len(triangle[i]); j++ {
			if i == 0 {
				triangle[i][j] = triangle[i][j]
			} else {
				if j == 0 {
					triangle[i][j] = triangle[i-1][j] + triangle[i][j]
				} else if j == len(triangle[i])-1 {
					triangle[i][j] = triangle[i-1][j-1] + triangle[i][j]
				} else {
					if triangle[i-1][j] < triangle[i-1][j-1] {
						triangle[i][j] = triangle[i-1][j] + triangle[i][j]
					} else {
						triangle[i][j] = triangle[i-1][j-1] + triangle[i][j]
					}
				}
			}
		}
	}
	res := triangle[len(triangle)-1][0]
	for _, v := range triangle[len(triangle)-1] {
		if v < res {
			res = v
		}
	}
	return res
}
func main() {
	s := [][]int{{2},
		{3, 4},
		{6, 5, 7},
		{4, 1, 8, 3}}
	log.Println(minimumTotal(s))
}

/*
状态方程
f(n,j) = min(f(n-1,j-1),f(n-1,j))+p[i,j]
// 转移矩阵
f(n,j) = min(f(n-1 ,j-1), f(n-1 ,j)) + p[n, j];
// f(n, j) 表示从1到n行的，以j为终点的最短路径。
// f(n-1, j) 表示从1到n-1行的，以j为终点的最短路径。
// p[n, j] 表示第n行的第j个元素
*/
