package main

func re(node *Node) *Node {
	new := re(node.Next)
	node.Next.Next = new
	node.Next = nil
	return node
}

func re1(node *Node) *Node {
	var preNode, nextNode *Node
	for node != nil {
		nextNode = node.Next
		node.Next = preNode
		preNode = node
		node = nextNode
	}
	return preNode
}
