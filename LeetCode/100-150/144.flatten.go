package _00_150

/*给定一个二叉树，原地将它展开为链表。

例如，给定二叉树

    1
   / \
  2   5
 / \   \
3   4   6
将其展开为：

1
 \
  2
   \
    3
     \
      4
       \
        5
         \
          6*/
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}



func flatten(root *TreeNode) {
	var pre *TreeNode
	if root==nil{
		return
	}
	flatten(root.Right)
	flatten(root.Left)
	root.Right = pre
	root.Left = nil
	pre = root
}
