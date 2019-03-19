package main

import (
	"container/list"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//         8
//        / \
//       3   9
//     /   \
//    2     4
//   / \   / \
//  1   5  6  7

///================================================递归方式
//中序遍历--递归方式
func midRange(root *TreeNode, res []int) []int {
	if root == nil {
		return res
	}
	res = midRange(root.Left, res)
	res = append(res, root.Val)
	res = midRange(root.Right, res)
	return res
}

//先序遍历--递归方式
func firstRange(root *TreeNode, res []int) []int {
	if root == nil {
		return res
	}
	res = append(res, root.Val)
	res = firstRange(root.Left, res)
	res = firstRange(root.Right, res)
	return res
}

//后序遍历--递归方式
func lastRange(root *TreeNode, res []int) []int {
	if root == nil {
		return res
	}
	res = lastRange(root.Left, res)
	res = lastRange(root.Right, res)
	res = append(res, root.Val)
	return res
}

///================================================非递归方式
// 中序遍历-非递归
func midRangeNoRecursion(t *TreeNode) []int {
	stack := list.New()
	res := make([]int, 0)
	for t != nil || stack.Len() != 0 {
		for t != nil {
			stack.PushBack(t)
			t = t.Left
		}
		if stack.Len() != 0 {
			v := stack.Back()
			res = append(res, t.Val) //visit
			t = t.Right
			stack.Remove(v)
		}
	}
	return res
}

// 先序遍历-非递归
func firsteRangeNoRecursion(t *TreeNode) []int {
	stack := list.New()
	res := make([]int, 0)
	for t != nil || stack.Len() != 0 {
		for t != nil {
			res = append(res, t.Val) //visit
			stack.PushBack(t)
			t = t.Left
		}
		if stack.Len() != 0 {
			v := stack.Back()
			t = t.Right
			stack.Remove(v)
		}
	}
	return res
}

// 后序遍历-非递归
func lastRangeNoRecursion(t *TreeNode) []int {
	stack := list.New()
	res := make([]int, 0)
	var preVisited *TreeNode
	for t != nil || stack.Len() != 0 {
		for t != nil {
			stack.PushBack(t)
			t = t.Left
		}
		v := stack.Back()
		top := v.Value.(*TreeNode)
		if (top.Left == nil && top.Right == nil) || (top.Right == nil && preVisited == top.Left) || preVisited == top.Right {
			res = append(res, top.Val) //visit
			preVisited = top
			stack.Remove(v)
		} else {
			t = top.Right
		}
	}
	return res
}

//info
/*
难理解的是非递归的额后序遍历
1、创建一个栈对象，根节点进栈，p赋初值null；

2、若栈非空，则栈顶结点的非空左孩子相继进栈；

3、若栈非空，查看栈顶结点，若栈顶结点的右孩子为空，或与p相等，则将栈顶结点弹出栈并访问它，同时将p指向该结点，并设置flag为true。否则将栈顶结点的右孩子（1个）压入栈中，并置flag为false

4、若flag==true,重复执行3，否则，重复执行步骤2和3，直到栈为空。
*/

///================================================知道其中两种求第三种
//ABDEFC
//DBFEAC
///================================================my
func getNode(fir, mid []int) *TreeNode {
	if len(fir) == 0 {
		return nil
	}
	node := &TreeNode{}
	node.Val = fir[0]
	index := myIndexOf(fir[0], mid)
	nextMidLeft := mid[0:index]
	nextMidRight := mid[index+1:]
	nextFirLeft := nextNextFir(fir, nextMidLeft)
	nextFirRight := nextNextFir(fir, nextMidRight)
	node.Left = getNode(nextFirLeft, nextMidLeft)
	node.Right = getNode(nextFirRight, nextMidRight)
	return node
}

func myIndexOf(val int, sli []int) int {
	for i := range sli {
		if sli[i] == val {
			return i
		}
	}
	return -1
}

func nextNextFir(ori, rest []int) []int {
	restMap := make(map[int]int)
	res := make([]int, 0)
	for _, v := range rest {
		restMap[v] = 1
	}
	for _, v := range ori {
		if rest[v] != 0 {
			res = append(res, v)
		}
	}
	return res
}

///================================================大佬
func buildTreeWithPreAndMid(pre, in []int) *TreeNode {

	if len(in) == 0 {
		return nil
	}

	res := &TreeNode{
		Val: pre[0],
	}

	if len(in) == 1 {
		return res
	}

	idx := indexOf(res.Val, in)

	res.Left = buildTreeWithPreAndMid(pre[1:idx+1], in[:idx])
	res.Right = buildTreeWithPreAndMid(pre[idx+1:], in[idx+1:])

	return res
}

func indexOf(val int, nums []int) int {
	for i, v := range nums {
		if v == val {
			return i
		}
	}
	return 0
}

func buildTreeWithPreAndAft(in []int, post []int) *TreeNode {
	if len(in) == 0 {
		return nil
	}

	res := &TreeNode{
		Val: post[len(post)-1],
	}

	if len(in) == 1 {
		return res
	}

	idx := indexOf(res.Val, in)

	res.Left = buildTreeWithPreAndAft(in[:idx], post[:idx])
	res.Right = buildTreeWithPreAndAft(in[idx+1:], post[idx:len(post)-1])

	return res
}
