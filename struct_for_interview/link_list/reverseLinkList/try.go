package main

func reverseSingleLinkList(node *Node) *Node {
	if node.Next == nil {
		return node
	}
	temp := reverseSingleLinkList(node)
	node.Next.Next = node
	node.Next = nil
	return temp
}

func reverseSingleLinkList2(head *Node) *Node {
	var prev, next *Node
	for head != nil {
		next = head.Next
		head.Next = prev
		prev = head
		head = next
	}
	return prev
}

//递归
func reverse5(node *DoubleNode) *DoubleNode {
	if node == nil || node.Next == nil {
		return node
	}
	newNode := reverse4(node.Next)
	node.Next.Next = node
	node.Next.Pre = newNode

	node.Pre = node.Next
	node.Next = nil
	return newNode
}
