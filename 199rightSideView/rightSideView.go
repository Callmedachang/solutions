package _99rightSideView

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func rightSideView(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	res := make([]int, 0)
	needList := []*TreeNode{root}
	for len(needList) > 0 {
		res = append(res, needList[len(needList)-1].Val)
		needList = getNextLevel(needList)
	}
	return res
}

func getNextLevel(nodes []*TreeNode) (res []*TreeNode) {
	for _, v := range nodes {
		if v.Left != nil {
			res = append(res, v.Left)
		}
		if v.Right != nil {
			res = append(res, v.Right)
		}
	}
	return
}
