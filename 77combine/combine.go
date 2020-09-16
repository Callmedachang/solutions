package _7combine

/*
给定两个整数 n 和 k，返回 1 …… n 中所有可能的 k 个数的组合。

示例:

输入 n = 4, k = 2
输出:
[
  [2,4],
  [3,4],
  [2,3],
  [1,2],
  [1,3],
  [1,4],
]
*/

func Combine(n int, k int) [][]int {
	if n == k {
		res := make([][]int, 1)
		for i := 1; i <= n; i++ {
			res[0] = append(res[0], i)
		}
		return res
	}
	res := Combine2(0, n, k)
	return res
}

func Combine2(start, n, k int) [][]int {
	res := make([][]int, 0)
	if k == 1 {
		for i := start + 1; i <= n; i++ {
			if start > 0 {
				res = append(res, []int{start, i})
			} else {
				res = append(res, []int{i})
			}
		}
		return res
	}
	for i := start; i <= n-k; i++ {
		tempRes := Combine2(i+1, n, k-1)
		res = append(res, tempRes...)
	}
	if start>0{
		for i, v := range res {
			res[i] = append([]int{start}, v...)
		}
	}
	return res
}
