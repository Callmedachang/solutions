package interview_25Binary_Tree_Trace

import "log"

/*
题目：输入一棵二叉树和一个整数，打印出二叉树中结点值的和为输入整数的所有路径。
从树的根结点开始往下一直到叶结点所经过的结点形成一条路径。
*/

type Node struct {
	Val   int
	Left  *Node
	Right *Node
}

func BinaryTreeTrace(root *Node, targetValue int) {
	RootFirstRange(root, []int{}, targetValue)
}

func RootFirstRange(root *Node, res []int, target int) []int {
	if root == nil {
		return res
	}
	res = append(res, root.Val)
	res = RootFirstRange(root.Left, res, target)
	res = RootFirstRange(root.Right, res, target)
	total := 0
	for i := 0; i < len(res); i++ {
		total += res[i]
	}
	if total == target {
		log.Println(res)
	}
	res = res[:len(res)-1]
	return res
}

//func RootFirstRange(root *Node) {
//	stack, res := list.New(), make([]int, 0)
//	for root != nil || stack.Len() > 0 {
//		for root != nil {
//			res = append(res, root.Val)
//			stack.PushBack(root)
//			root=root.Left
//		}
//		if stack.Len() > 0 {
//			v := stack.Back()
//			s:= v.Value.(*Node)
//			root = s.Right
//			stack.Remove(v)
//			res = res[:len(res)-1]
//		}
//	}
//}
