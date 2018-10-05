package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	head = reverseLink(head)
	before := head
	next := before.Next
	for i := 2; i < n; i++ {
		before = before.Next
		next = next.Next
	}
	if n==1{
		head=next
	}else{
		before.Next = next.Next
	}
	return reverseLink(head)
}
func reverseLink(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	var reversedHead *ListNode
	var p *ListNode
	reversedHead = head
	head = head.Next
	reversedHead.Next = nil
	p = head.Next
	for head != nil {
		head.Next = reversedHead
		reversedHead = head
		head = p
		if p != nil {
			p = p.Next
		}
	}
	return reversedHead
}
func main() {
	ss := &ListNode{Val: 1, Next: &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, nil}}}}}
	ss = removeNthFromEnd(ss, 2)
	fmt.Println(ss)
}
