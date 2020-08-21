package interview_39DeepOfTree
/*
：输入一棵二叉树的根结点，求该树的深度。
从根结点到叶结点依次经过的结点（含根、叶结点）形成树的一条路径，最长路径的长度为树的深度。
*/
type Node struct {
	Val   int
	Left  *Node
	Right *Node
}

func DeepOfTree(root *Node) int {
	if root == nil {
		return 0
	}
	return max(DeepOfTree(root.Left), DeepOfTree(root.Right)) + 1
}
func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
