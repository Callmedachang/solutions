package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func GetValue(linkLis *ListNode, index int) int {
	this := linkLis
	if index == 0 {
		return this.Val
	}
	for i := 0; i < index; i++ {
		if this == nil {
			return 0
		}
		this = this.Next
	}
	if this == nil {
		return 0
	}
	return this.Val
}
func GetNode(linkLis *ListNode, index int) *ListNode {
	this := linkLis
	if index == 0 {
		return this
	}
	for i := 0; i < index; i++ {
		this = this.Next
	}
	return this
}
func Len(linkLis *ListNode) int {
	this := linkLis
	len := 0
	for {
		if this != nil {
			this = this.Next
			len++
		} else {
			break
		}
	}
	return len
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	len := 0
	res := &ListNode{}
	if Len(l1) > Len(l2) {
		len = Len(l1)
	} else {
		len = Len(l2)
	}
	addIndex := 0
	for i := 0; i < len; i++ {
		addNum := 0
		if GetValue(l1, i)+GetValue(l2, i)+addIndex > 9 {
			addNum = GetValue(l1, i) + GetValue(l2, i) + addIndex - 10
			addIndex = 1
		} else {
			addNum = GetValue(l1, i) + GetValue(l2, i) + addIndex
			addIndex = 0
		}
		GetNode(res, i).Val = addNum
		if i != len-1 {
			GetNode(res, i).Next = &ListNode{}
		}
	}
	if addIndex == 1 {
		GetNode(res, Len(res)-1).Next = &ListNode{1, nil}
	}
	return res
}
