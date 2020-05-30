package resove

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return isSymm(root.Left, root.Right)
}

func isSymm(n1, n2 *TreeNode) bool {
	if n1 == nil && n2 == nil {
		return true
	}
	if (n1 == nil && n2 != nil) || (n2 == nil && n1 != nil) {
		return false
	}
	if isSymm(n1.Left, n2.Right) && isSymm(n1.Right, n2.Left) && (n1.Val == n2.Val) {
		return true
	}
	return false
}
