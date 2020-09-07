package _21maximalSquare

/*
在一个由 0 和 1 组成的二维矩阵内，找到只包含 1 的最大正方形，并返回其面积。

示例:

输入:

1 0 1 0 0
1 0 1 1 1
1 1 1 1 1
1 0 0 1 0

输出: 4
*/
//状态转移方程：dp(i,j)=min(dp(i−1,j),dp(i−1,j−1),dp(i,j−1))+1
func MaximalSquare(matrix [][]byte) int {
	max := 0
	dp := make([][]int, len(matrix))
	for i := range dp {
		dp[i] = make([]int, len(matrix[0]))
	}
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			if matrix[i][j] == '1' {
				dp[i][j] = next(dp, i, j)
				if dp[i][j] > max {
					max = dp[i][j]
				}
			}
		}
	}
	return max * max
}

func next(dp [][]int, i, j int) int {
	if j == 0 || i == 0 {
		return 1
	}
	if dp[i-1][j-1] <= dp[i][j-1] && dp[i-1][j-1] <= dp[i-1][j] {
		return dp[i-1][j-1] + 1
	}
	if dp[i-1][j] <= dp[i-1][j-1] && dp[i-1][j] <= dp[i][j-1] {
		return dp[i-1][j] + 1
	}
	return dp[i][j-1] + 1
}
