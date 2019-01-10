package _0MaxDepth

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	res := 0
	nodes := make([]*TreeNode, 0)
	nodes = append(nodes, root)
	for len(nodes) != 0 {
		res++
		temp := make([]*TreeNode, 0)
		for _, v := range nodes {
			if v.Left != nil {
				temp = append(temp, v.Left)
			}
			if v.Right != nil {
				temp = append(temp, v.Right)
			}
		}
		nodes = temp
	}
	return res
}
