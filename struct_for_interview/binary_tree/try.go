package main

import (
	"container/list"
	"log"
)

type TreeNode2 struct {
	Val   int
	Left  *TreeNode2
	Right *TreeNode2
}

//二叉树的各种遍历方式

//二叉树的中序遍历 递归与非递归
func midRootRange(root *TreeNode2) []int {
	return midRootRangeImpl(root, []int{})
}

func midRootRangeImpl(root *TreeNode2, res []int) []int {
	if root == nil {
		return res
	}
	res = midRootRangeImpl(root.Left, res)
	res = append(res, root.Val)
	res = midRootRangeImpl(root.Right, res)
	return res
}

func midRootRangeImplNotUseRecursion(root *TreeNode2) []int {
	stack, res := list.New(), make([]int, 0)
	for root != nil || stack.Len() > 0 {
		for root != nil {
			stack.PushBack(root)
			root = root.Left
		}
		if stack.Len() > 0 {
			v := stack.Back()
			s := v.Value.(*TreeNode2)
			res = append(res, s.Val)
			root = s.Right
			stack.Remove(v)
		}

	}
	return res
}

//二叉树先序遍历的递归与非递归
func firstRootRange(root *TreeNode2) []int {
	return firstRootRangeImpl(root, []int{})
}

func firstRootRangeImpl(root *TreeNode2, res []int) []int {
	if root == nil {
		return res
	}
	res = append(res, root.Val)
	res = firstRootRangeImpl(root.Left, res)
	res = firstRootRangeImpl(root.Right, res)
	return res
}

func firstRootImplWithoutRecursion(root *TreeNode2) []int {
	stack, res := list.New(), make([]int, 0)
	for root != nil || stack.Len() > 0 {
		for root != nil {
			res = append(res, root.Val)
			stack.PushBack(root)
			root = root.Left
		}
		if stack.Len() > 0 {
			v := stack.Back()
			s := v.Value.(*TreeNode2)
			root = s.Right
			stack.Remove(v)

		}
	}
	return res
}

//二叉树后序遍历的递归与非递归

func lastRootRange(root *TreeNode2) []int {
	return lastRootRangeImpl(root, []int{})
}

func lastRootRangeImpl(root *TreeNode2, res []int) []int {
	if root == nil {
		return res
	}
	res = firstRootRangeImpl(root.Left, res)
	res = firstRootRangeImpl(root.Right, res)
	res = append(res, root.Val)
	return res
}

func lastRootImplWithoutRecursion(root *TreeNode2) []int {
	stack, res := list.New(), make([]int, 0)
	preNode := &TreeNode2{}
	for root != nil || stack.Len() > 0 {
		for root != nil {
			stack.PushBack(root)
			root = root.Left
		}
		current := stack.Back()
		currentNode := current.Value.(*TreeNode2)
		if (currentNode.Left == nil && currentNode.Right == nil) || (currentNode.Left == preNode && currentNode.Right == nil) || (currentNode.Right == preNode) {
			preNode = currentNode
			res = append(res, currentNode.Val)
			stack.Remove(current)
		} else {
			root = currentNode.Right
		}
	}
	return res
}

//重建二叉树

func buildWithFirstAndMid(fir, mid []int) *TreeNode2 {
	if len(fir) == 0 {
		return nil
	}
	root := &TreeNode2{Val: fir[0]}
	idx := indexOf(root.Val, mid, )
	root.Left = buildWithFirstAndMid(fir[1:idx+1], mid[:idx])
	root.Right = buildWithFirstAndMid(fir[idx+1:], mid[idx+1:])
	return root
}

func buildWithMidAndLast(mid, last []int) *TreeNode2 {
	if len(last) == 0 {
		return nil
	}
	root := &TreeNode2{Val: last[len(last)-1]}
	idx := indexOf(root.Val, mid)
	root.Left = buildWithMidAndLast(mid[:idx], last[:idx])
	root.Right = buildWithMidAndLast(mid[idx:], last[idx:len(last)-1])
	return root
}

func main() {
	t := &TreeNode2{Val: 1, Left: &TreeNode2{Val: 2}, Right: &TreeNode2{Val: 3}}
	log.Println(firstRootRange(t))
	log.Println(lastRootImplWithoutRecursion(t))
}
