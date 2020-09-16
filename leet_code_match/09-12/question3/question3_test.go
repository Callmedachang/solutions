package question3

import (
	"container/list"
	"log"
	"testing"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func hasPathSum(root *TreeNode, sum int) bool {
	if root==nil{
		return false
	}
	stack := list.New()
	stack.PushBack(root)
	for stack.Len() != 0 {
		s := stack.Back()
		stack.Remove(s)
		node, _ := s.Value.(*TreeNode)
		if node.Left == nil && node.Right == nil {
			if node.Val == sum {
				return true
			}
		}
		if node.Left != nil {
			node.Left.Val += node.Val
			stack.PushBack(node.Left)
		}
		if node.Right != nil {
			node.Right.Val += node.Val
			stack.PushBack(node.Right)
		}
	}
	return false
}

func TestMinCostConnectPoints(t *testing.T) {
	//s:=[][]int{{0,0},{2,2},{3,10},{5,2},{7,0}}
	//s:=[][]int{{0,0},{2,2},{3,10},{5,2},{7,0}}
	s:=[][]int{{-14,-14},{-18,5},{18,-10},{18,18},{10,-2}}
	log.Println(MinCostConnectPoints(s))
}