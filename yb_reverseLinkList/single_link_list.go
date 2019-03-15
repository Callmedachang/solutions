package main

import "log"

type Node struct {
	Val  int
	Next *Node
}

//迭代
func reverse1(head *Node) *Node {
	var preNode, nextNode *Node
	for head != nil {
		//  保存头节点的下一个节点，
		nextNode = head.Next
		//  将头节点指向前一个节点
		head.Next = preNode
		//  更新前一个节点
		preNode = head
		//  更新头节点
		head = nextNode
	}
	return preNode
}

//递归
func reverse2(node *Node) *Node {
	if  node.Next == nil {
		return node
	}
	var newNode = reverse2(node.Next)
	// 将第二个节点的指针指向头节点
	node.Next.Next = node
	// 将头几点的指针置为空
	node.Next = nil
	return newNode
}

func main() {
	node1 := new(Node)
	node1.Val = 1
	node2 := new(Node)
	node2.Val = 2
	node3 := new(Node)
	node3.Val = 3
	node4 := new(Node)
	node4.Val = 4
	node1.Next = node2
	node2.Next = node3
	node3.Next = node4
	s := reverse1(node1)
	s = reverse2(s)
	log.Println(s)
}
