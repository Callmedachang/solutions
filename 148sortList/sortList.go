package main

import "log"

type ListNode struct {
	Val  int
	Next *ListNode
}

func sortList(head *ListNode) *ListNode {
	needTurn := true
	for needTurn {
		indexNode := head
		for indexNode.Next != nil {
			needTurn = false
			if indexNode.Val > indexNode.Next.Val {
				needTurn = true
				temp := indexNode.Next
				indexNode.Next = indexNode.Next.Next
				temp.Next = indexNode
				continue
			} else {
				indexNode = indexNode.Next
			}
		}
	}
	return head
}
func main() {
	n := &ListNode{Val: 2, Next: &ListNode{Val: 1}}
	b := sortList(n)
	log.Println(b)
}
