package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func minDepth(root *TreeNode) int {
	if root != nil {
		return 1 + min(minDepth(root.Right), minDepth(root.Left))
	} else {
		return 0
	}
}
func min(q, b int) int {
	if q < b {
		return q
	}
	return b
}
