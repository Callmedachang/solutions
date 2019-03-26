package _5generateTrees

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func generateTrees(n int) []*TreeNode {

}
func s(trees []*TreeNode, n int) []*TreeNode {
	if n < 1 {
		return trees
	}
	if len(trees) == 0 {
		trees = append(trees, &TreeNode{Val: n})
	}else{

	}

}
