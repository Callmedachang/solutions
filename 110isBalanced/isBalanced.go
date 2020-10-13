package _10isBalanced

/*
给定一个二叉树，判断它是否是高度平衡的二叉树。

本题中，一棵高度平衡二叉树定义为：

一个二叉树每个节点 的左右两个子树的高度差的绝对值不超过 1。

示例 1:

给定二叉树 [3,9,20,null,null,15,7]

    3
   / \
  9  20
    /  \
   15   7
返回 true 。

示例 2:

给定二叉树 [1,2,2,3,3,null,null,4,4]

       1
      / \
     2   2
    / \
   3   3
  / \
 4   4
返回 false 。

*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isBalanced(root *TreeNode) bool {
	res, _ := isBalancedInner(root)
	return res
}

func isBalancedInner(root *TreeNode) (bool, int) {
	if root == nil {
		return true, 0
	}
	isLeftBalance, leftHeight := isBalancedInner(root.Left)
	if !isLeftBalance {
		return false, 0
	}
	isRightBalance, rightHeight := isBalancedInner(root.Right)
	if !isRightBalance {
		return false, 0
	}

	if max(leftHeight, rightHeight)-leftHeight > 1 || max(leftHeight, rightHeight)-rightHeight > 1 {
		return false, 0
	}
	return true, max(leftHeight, rightHeight) + 1
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
