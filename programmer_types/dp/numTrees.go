package main

/*
给定一个整数 n，求以 1 …… n 为节点组成的二叉搜索树有多少种？

示例:

输入: 3
输出: 5
解释:
给定 n = 3, 一共有 5 种不同结构的二叉搜索树:

   1         3     3      2      1
    \       /     /      / \      \
     3     2     1      1   3      2
    /     /       \                 \
   2     1         2                 3

链接：https://leetcode-cn.com/problems/unique-binary-search-trees
*/

//递归思路
func numTrees(n int) int {
	res := 0
	list := make([]int, 0)
	for i := 1; i <= n; i++ {
		list = append(list, i)
	}
	for i := 1; i <= n; i++ {
		res += find(i, removeIndex(i-1, list))
	}
	return res
}

func find(root int, target []int) int {
	if len(target) <= 1 {
		return 1
	}
	left, right := make([]int, 0), make([]int, 0)
	for _, v := range target {
		if v < root {
			left = append(left, v)
		} else {
			right = append(right, v)
		}
	}
	leftNum, rightNum := 0, 0
	for index, v := range left {
		leftNum += find(v, removeIndex(index, left))
	}
	for index, v := range right {
		rightNum += find(v, removeIndex(index, right))
	}
	if leftNum == 0 {
		leftNum = 1
	}
	if rightNum == 0 {
		rightNum = 1
	}
	return leftNum * rightNum
}

func removeIndex(i int, src []int) []int {
	res := make([]int, len(src))
	trueRes := make([]int, 0)
	for index := range src {
		res[index] = src[index]
	}
	trueRes = append(res[0:i], res[i+1:]...)
	return trueRes
}

//动态规划思路
/*
思路:

动态规划

假设 n 个节点存在二叉排序树的个数是 G(n)，令 f(i)为以 i 为根的二叉搜索树的个数

即有:G(n) = f(1) + f(2) + f(3) + f(4) + …… + f(n)

n 为根节点，当 i 为根节点时，其左子树节点个数为[1,2,3,……,i-1]，右子树节点个数为[i+1,i+2,……n]，所以当 i 为根节点时，其左子树节点个数为 i-1 个，右子树节点为 n-i，即 f(i) = G(i-1)*G(n-i),

上面两式可得:G(n) = G(0)*G(n-1)+G(1)*(n-2)+……+G(n-1)*G(0)
*/

func numTrees2(n int) int {
	if n == 1 || n == 0 {
		return 1
	}
	res := 0
	for i := 0; i < n; i++ {
		res += numTrees2(i) * numTrees2(n-1-i)
	}
	return res
}

//继续优化从 1，2，3分别开始算,这样可以在求N的时候知道之前的n-1的值 进一步加快速度
func numTrees3(n int)int{
	dp := make([]int,n+1)
	dp[0], dp[1] = 1, 1
	for i := 2; i <= n; i++ {
		for j := 1; j <= i ; j++ {
			dp[i] += dp[j-1]*dp[i-j]
		}
	}
	return dp[n]
}