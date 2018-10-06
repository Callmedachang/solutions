package main

import "fmt"

/**
 * Definition for singly-linked list.*/
type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeKLists(lists []*ListNode) *ListNode {
	listsNUms := make([]int, 0)
	var res *ListNode
	if len(lists) == 0 {
		return nil
	}
	pic := res
	resLen := 0
	lenTemp := lists
	len := 0
	for _, v := range lenTemp {
		for v != nil {
			len++
			v = v.Next
		}
	}
	for resLen < len {
		tempMin := 99999999999
		tempIndex := 0
		for i, v := range lists {
			if v != nil {
				if v.Val < tempMin {
					tempMin = v.Val
					tempIndex = i
				}

			}
		}
		listsNUms = append(listsNUms, tempMin)
		resLen++
		lists[tempIndex] = lists[tempIndex].Next
	}
	for i, v := range listsNUms {
		if i == 0 {
			pic.Val = v
		} else {
			pic.Next = &ListNode{Val: v, Next: nil}
			pic = pic.Next
		}
	}
	return res
}
func main() {
	l1 := &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{Val: 3, Next: &ListNode{Val: 4}}}}}
	l2 := &ListNode{5, &ListNode{6, &ListNode{7, nil}}}
	l3 := &ListNode{7, &ListNode{8, &ListNode{9, nil}}}
	in := make([]*ListNode, 0)
	in = append(in, l1, l2, l3)
	ss := mergeKLists(in)
	fmt.Println(ss)
}
