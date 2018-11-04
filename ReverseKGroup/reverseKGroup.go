package main

import "fmt"

/*
给出一个链表，每 k 个节点一组进行翻转，并返回翻转后的链表。

k 是一个正整数，它的值小于或等于链表的长度。如果节点总数不是 k 的整数倍，那么将最后剩余节点保持原有顺序。

示例 :

给定这个链表：1->2->3->4->5

当 k = 2 时，应当返回: 2->1->4->3->5

当 k = 3 时，应当返回: 3->2->1->4->5

说明 :

你的算法只能使用常数的额外空间。
你不能只是单纯的改变节点内部的值，而是需要实际的进行节点交换*/
type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	headCopy := head
	if headCopy == nil {
		return head
	}
	for i := 1; i < k; i++ {
		if headCopy.Next == nil {
			return head
		} else {
			headCopy = headCopy.Next
		}
	}
	if head == nil || head.Next == nil {
		return head
	}
	tailNext := headCopy.Next
	headCopy.Next=nil
	head, headCopy = reverse(head, headCopy)
	headCopy.Next = reverseKGroup(tailNext, k)
	return head
}
func reverse(head, tail *ListNode) (*ListNode, *ListNode) {
	curPre, cur := head, head.Next
	for cur != nil {
		curPre, cur, cur.Next = cur, cur.Next, curPre
	}
	return tail, head
}

func main() {
	l1 := &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, nil}}}}
	l2 := reverseKGroup(l1, 3)
	fmt.Println(l2)
}
