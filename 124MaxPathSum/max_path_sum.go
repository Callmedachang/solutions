package main

import "log"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var maxSum = -1 << 32

func maxPathSum(root *TreeNode) int {
	getMaxSum(root)
	return maxSum
}

func getMaxSum(node *TreeNode) int {
	if node == nil {
		return 0
	}
	left := max(0, getMaxSum(node.Left))
	right := max(0, getMaxSum(node.Right))
	maxSum = max(maxSum, node.Val+left+right)
	return node.Val + max(left, right)

}
func max(v1, v2 int) int {
	if v1 > v2 {
		return v1
	} else {
		return v2
	}
}

func main() {
	log.Print(maxPathSum(&TreeNode{Val: 0}))
}
