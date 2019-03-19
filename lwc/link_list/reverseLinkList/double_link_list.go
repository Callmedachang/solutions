package main

type DoubleNode struct {
	Val  int
	Pre  *DoubleNode
	Next *DoubleNode
}

//迭代
func reverse3(head *DoubleNode) *DoubleNode {
	// 先声明两个节点，并将两个节点都置为空，这里将两个新声明的节点都置为空，才能保证反转后的链表，头节点的前指针为空，尾节点的后指针为空
	PreNode := new(DoubleNode)
	// head为头节点，也就是当前节点
	for head != nil {
		// 保存第二个节点的值
		NextNode := head.Next
		// 将当前节点的前指针，指向后一个节点，也就是把头节点的后节点变成前节点，
		head.Pre = NextNode
		// 将当前节点的后指针，指向前一个节点，也就是把头节点的前节点变成后节点。
		head.Next = PreNode
		// 更新前节点，也就是把头节点变成前节点
		PreNode = head
		// 更新当前节点，也就是把后节点变成头节点
		head = NextNode
	}
	return PreNode
}

//递归
func reverse4(node *DoubleNode) *DoubleNode {
	if node == nil || node.Next == nil {
		return node
	}
	var newNode = reverse4(node.Next)
	// 将第二个节点的指针指向头节点
	node.Next.Next = node
	node.Next.Pre = newNode
	// 将头几点的指针置为空
	node.Pre = node.Next
	node.Next = nil
	return newNode
}

/*
mark
非递归的实现：
用一个pre来记录遍地到N的每一个节点之前已经反转好的链表，到最后preNode就是整个被反转的链表:preNode+Head
递归的实现：
递归点：
假设这个节点之后所有节点已经反转OK
递归返回点：
当前节点为空或者next为nil
*/
