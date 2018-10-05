package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	var res *ListNode
	if l1==nil{
		return l2
	}
	if l2== nil{
		return l1
	}
	for l1 != nil && l2 != nil {
		if l1.Val > l2.Val {
			if res == nil {
				res = &ListNode{l2.Val, nil}
			} else {
				addNode(res,l2.Val)
			}
			l2 = l2.Next
		} else {
			if res == nil {
				res = &ListNode{l1.Val, nil}
			} else {
				addNode(res,l1.Val)
			}
			l1 = l1.Next
		}
	}
	if l1 == nil && l2 != nil {
		for l2 != nil {
			addNode(res,l2.Val)
			l2 = l2.Next
		}
	}
	if l1 != nil && l2 == nil {
		for l1 != nil {
			addNode(res,l1.Val)
			l1 = l1.Next
		}
	}
	return res
}
func addNode(l1 *ListNode, val int) {
	temp := l1
	for temp.Next != nil {
		temp = temp.Next
	}
	temp.Next = &ListNode{Val: val, Next: nil}
}
func main() {
	l1 := &ListNode{1, &ListNode{2, &ListNode{3, nil}}}
	l2 := &ListNode{1, &ListNode{3, &ListNode{4, nil}}}
	l3 := mergeTwoLists(l1, l2)
	fmt.Println(l3)
}
