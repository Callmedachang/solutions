package main

import "log"

/**
给出两个 非空 的链表用来表示两个非负的整数。其中，它们各自的位数是按照 逆序 的方式存储的，并且它们的每个节点只能存储 一位 数字。

如果，我们将这两个数相加起来，则会返回一个新的链表来表示它们的和。

您可以假设除了数字 0 之外，这两个数都不会以 0 开头。

示例：

输入：(2 -> 4 -> 3) + (5 -> 6 -> 4)
输出：7 -> 0 -> 8
原因：342 + 465 = 807/
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	res := &ListNode{}
	temp := res
	up := 0
	for l1 != nil || l2 != nil || up > 0 {
		v1, v2 := 0, 0
		if l1 != nil {
			v1 = l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			v2 = l2.Val
			l2 = l2.Next
		}
		cu := v1 + v2 + up
		up = 0
		for cu > 9 {
			cu = cu - 10
			up++
		}
		temp.Val = cu
		if l1 != nil || l2 != nil  || up > 0 {
			temp.Next = &ListNode{}
			temp = temp.Next
		}
	}
	return res
}
func main() {
	s:=addTwoNumbers(&ListNode{1, &ListNode{8,nil}}, &ListNode{0, nil})
	log.Println(s)
}
