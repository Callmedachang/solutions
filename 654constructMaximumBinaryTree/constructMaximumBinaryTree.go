package main

/*
给定一个不含重复元素的整数数组。一个以此数组构建的最大二叉树定义如下：

二叉树的根是数组中的最大元素。
左子树是通过数组中最大值左边部分构造出的最大二叉树。
右子树是通过数组中最大值右边部分构造出的最大二叉树。
通过给定的数组构建最大二叉树，并且输出这个树的根节点。

Example 1:

输入: [3,2,1,6,0,5]
输入: 返回下面这棵树的根节点：

      6
    /   \
   3     5
    \    /
     2  0
       \
        1
注意:

给定的数组的大小在 [1, 1000] 之间。
*/
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func constructMaximumBinaryTree(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}
	val, index := getMax(nums)
	return &TreeNode{Val: val, Left: constructMaximumBinaryTree(nums[0:index]), Right: constructMaximumBinaryTree(nums[index+1:])}
}

func getMax(a []int) (int, int) {
	res, index := 0, 0
	for i, k := range a {
		if k > res {
			res = k
			index = i
		}
	}
	return res, index
}

func main() {
	constructMaximumBinaryTree([]int{3, 2, 1, 6, 0, 5})
}
