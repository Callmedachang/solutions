package main

import (
	"log"
)

/*

请判断一个链表是否为回文链表。

示例 1:

输入: 1->2
输出: false
示例 2:

输入: 1->2->2->1
输出: true
进阶：
你能否用 O(n) 时间复杂度和 O(1) 空间复杂度解决此题？
*/
func main() {
	//head := &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3,Next: &ListNode{Val: 2,Next: &ListNode{Val:1 }}}}}
	head := &ListNode{Val: 1, Next: &ListNode{Val: 1}}
	log.Println(isPalindrome2(head))
	//head := &ListNode{Val: 1}
	//log.Println(isPalindrome(head))
}

type ListNode struct {
	Val  int
	Next *ListNode
}

//1  2  3  2  1
func isPalindrome(head *ListNode) bool {
	list := make([]int, 0)
	for head != nil {
		list = append(list, head.Val)
		head = head.Next
	}
	listHead, listTail := 0, len(list)-1
	for listHead < listTail {
		if list[listHead] == list[listTail] {
			listHead++
			listTail--
		} else {
			return false
		}
	}
	return true
}

func isPalindrome2(head *ListNode) bool {
	if head==nil||head.Next==nil{
		return true
	}
	if head.Next.Next==nil{
		return head.Val==head.Next.Val
	}
	index, slow, fast := 0, head, head
	var next *ListNode
	for fast.Next != nil && fast.Next.Next != nil {
		if index == 0 {
			fast = fast.Next
		} else {
			slow = slow.Next
			fast = fast.Next.Next
		}
		index++
	}
	if fast.Next != nil {
		next = slow.Next.Next
	} else {
		next = slow.Next
	}
	slow.Next = nil
	prev := reverseList(head)
	for prev != nil {
		if prev.Val != next.Val {
			return false
		}
		prev, next = prev.Next, next.Next
	}
	return true
}

func reverseList(head *ListNode) *ListNode {
	var prev, next *ListNode
	for head != nil {
		next = head.Next
		head.Next = prev
		prev = head
		head = next
	}
	return prev
}

//func reverse1(head *Node) *Node {
//	var preNode, nextNode *Node
//	for head != nil {
//		//  保存头节点的下一个节点，
//		nextNode = head.Next
//		//  将头节点指向前一个节点
//		head.Next = preNode
//		//  更新前一个节点
//		preNode = head
//		//  更新头节点
//		head = nextNode
//	}
//	return preNode
//}
