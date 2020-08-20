package interview_27_BST_To_DoubleLinkList

type Node struct {
	Val   int
	Left  *Node
	Right *Node
}

func BSTToDoubleLinkList(root *Node) *Node {
	var last *Node
	head := Convert(root, last)
	for head.Left != nil {
		head = head.Left
	}
	return head
}
func Convert(root *Node, beforeLast *Node) *Node {
	if root == nil {
		return nil
	}
	currentNode := root
	if currentNode.Left != nil {
		beforeLast = Convert(currentNode.Left, beforeLast)
	}
	currentNode.Left = beforeLast
	if beforeLast != nil {
		beforeLast.Right = currentNode
	}
	beforeLast = currentNode
	if currentNode.Right!=nil{
		beforeLast=Convert(root.Right, beforeLast)
	}
	return beforeLast
}