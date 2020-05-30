package main

import "container/list"

func midRange2(root *TreeNode, res []int) []int {
	if root == nil {
		return res
	}
	res = midRange2(root.Left, res)
	res = append(res, root.Val)
	res = midRange2(root.Right, res)
	return res
}

func midRange2NotResursive(root *TreeNode) []int {
	list := list.New()
	res := make([]int, 0)
	for root != nil || list.Len() != 0 {
		for root != nil {
			list.PushBack(root)
			root = root.Left
		}
		if list.Len() > 0 {
			v := list.Back()
			node := v.Value.(*TreeNode)
			res = append(res, node.Val)
			if node.Right != nil {
				root = node.Right
			}
			list.Remove(v)
		}
	}
	return res
}

func preRange2(root *TreeNode, res []int) []int {
	if root == nil {
		return res
	}
	res = append(res, root.Val)
	res = preRange2(root.Left, res)
	res = preRange2(root.Right, res)
	return res
}

func preRange2NotResursive(root *TreeNode) []int {
	list := list.New()
	res := make([]int, 0)
	for list != nil || root != nil {
		for root != nil {
			res = append(res, root.Val)
			list.PushBack(root)
			root = root.Left
		}
		if list != nil {
			v := list.Back()
			root = v.Value.(*TreeNode).Right
			list.Remove(v)
		}
	}
	return res
}

func afterRange2(root *TreeNode, res []int) []int {
	if root == nil {
		return res
	}
	res = afterRange2(root.Left, res)
	res = afterRange2(root.Right, res)
	res = append(res, root.Val)
	return res
}

func afterRange2NotResursive(root *TreeNode) []int {
	list := list.New()
	res := make([]int, 0)
	var preNode *TreeNode
	for list != nil || root != nil {
		for root != nil {
			res = append(res, root.Val)
			list.PushBack(root)
			root = root.Left
		}
		v := list.Back()
		node := v.Value.(*TreeNode)
		if (node.Left == nil && node.Right == nil) || (node.Left == preNode && node.Right == nil) || (node.Right == preNode) {
			res = append(res, node.Val)
			preNode = node
			list.Remove(v)
		} else {
			root = node.Right
		}
	}
	return res
}
