package _1n_queen

import "log"

/*
n 皇后问题研究的是如何将 n 个皇后放置在 n×n 的棋盘上，并且使皇后彼此之间不能相互攻击。

给定一个整数 n，返回所有不同的 n 皇后问题的解决方案。

每一种解法包含一个明确的 n 皇后问题的棋子放置方案，该方案中 'Q' 和 '.' 分别代表了皇后和空位。

示例：
输入：4
输出：[
 [".Q……",  // 解法 1
  "……Q",
  "Q……",
  "……Q."],

 ["……Q.",  // 解法 2
  "Q……",
  "……Q",
  ".Q……"]
]
解释: 4 皇后问题存在两个不同的解法。
*/

func SolveNQueens(n int) {
	res := make([][]bool, n)
	for i := 0; i < n; i++ {
		res[i] = make([]bool, n)
	}
	find(res, 0, n)
}

func find(visit [][]bool, index, n int) bool {
	for i := 0; i < len(visit[index]); i++ {
		if visit[index][i] == false {
			if index == n-1 {
				visit[index][i] = true
				log.Println(visit)
				return true
			}
			visit[index][i] = true
			for j := index + 1; j < n; j++ {
				visit[j][i]= true
				if i+j-index > 0 && i+j-index < n {
					visit[j][i+j-index] = true
				}
				if i-j+index > 0 && i-j+index < n {
					visit[j][i-j+index] = true
				}
			}
			has := find(visit, index+1, n)
			visit[index][i] = false
			for j := index + 1; j < n; j++ {
				visit[j][i] = false
				if i+j-index > 0 && i+j-index < n {
					visit[j][i+j-index] = false
				}
				if i-j+index > 0 && i-j+index < n {
					visit[j][i-j+index] = false
				}
			}
			if has &&index>0{
				return has
			} else {
				continue
			}
		}
	}
	return false
}
