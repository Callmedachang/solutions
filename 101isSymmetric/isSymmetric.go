package _01isSymmetric

/*
给定一个二叉树，检查它是否是镜像对称的。

 

例如，二叉树 [1,2,2,3,4,4,3] 是对称的。

    1
   / \
  2   2
 / \ / \
3  4 4  3
 

但是下面这个 [1,2,2,null,3,null,3] 则不是镜像对称的:

    1
   / \
  2   2
   \   \
   3    3
 

进阶：

你可以运用递归和迭代两种方法解决这个问题吗？
*/
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return false
	}
	return Symmetric(root.Left, root.Right)
}
func Symmetric(p, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}
	if (p != nil && q == nil) || (q != nil && p == nil) {
		return false
	}
	if p.Val == q.Val {
		return Symmetric(p.Left, q.Right) && Symmetric(p.Right, q.Left)
	}
	return false
}
