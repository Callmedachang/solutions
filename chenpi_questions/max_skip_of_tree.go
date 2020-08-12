package main

type TreeNode struct {
	Val   interface{}
	Left  *TreeNode
	Right *TreeNode
}

func maxSkip(root *TreeNode) int {
	if root == nil {
		return 0
	}
	th := getMaxDeepOfTree(root.Left) + getMaxDeepOfTree(root.Left)
	return getMax([]int{maxSkip(root.Left), maxSkip(root.Right), th})
}

func getMaxDeepOfTree(root *TreeNode) int {
	if root == nil {
		return 0
	}
	left := getMaxDeepOfTree(root.Left)
	right := getMaxDeepOfTree(root.Right)
	if left > right {
		return left + 1
	}
	return right + 1
}

func getMax(x []int) int {
	if len(x) == 0 {
		return 0
	}
	res := x[0]
	for _, v := range x {
		if v > res {
			res = v
		}
	}
	return res
}
