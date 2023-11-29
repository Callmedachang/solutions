package _385TImeToCoverAllTrr

/*
给你一棵二叉树的根节点 root ，二叉树中节点的值 互不相同 。另给你一个整数 start 。在第 0 分钟，感染 将会从值为 start 的节点开始爆发。

每分钟，如果节点满足以下全部条件，就会被感染：
	1. 节点此前还没有感染。
	2. 节点与一个已感染节点相邻。
返回感染整棵树需要的分钟数。
*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func amountOfTime(root *TreeNode, start int) int {

}

func maxDeep(node *TreeNode) int64 {
	if node == nil {
		return 0
	}
	return max(maxDeep(node.Left), maxDeep(node.Left)) + 1
}

func findNode(root *TreeNode, start int) (int64, *TreeNode) {

	return 0, nil
}

func max(a, b int64) int64 {
	if a >= b {
		return a
	}
	return b

}
