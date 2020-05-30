package _00_150

import "log"


func sortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	s, q, pre := head, head, &ListNode{}
	for q != nil && q.Next != nil {
		pre = s
		s = s.Next
		q = q.Next.Next
	}
	pre.Next = nil
	left, right := head, s
	left = sortList(left)
	right = sortList(right)
	return mergeList(left, right)
}
func mergeList(l, r *ListNode) *ListNode {
	head := &ListNode{}
	index := head
	for l != nil && r != nil {
		index.Next = &ListNode{}
		index = index.Next
		if l.Val < r.Val {
			index.Val = l.Val
			l = l.Next
		} else {
			index.Val = r.Val
			r = r.Next
		}
	}
	if l != nil {
		index.Next = l
	}
	if r != nil {
		index.Next = r
	}
	return head.Next
}

func main() {
	s := sortList(&ListNode{4, &ListNode{2, &ListNode{3, &ListNode{1, nil}}}})
	log.Println(s)
	//s:=mergeList(&ListNode{1, &ListNode{2,nil}},&ListNode{3, &ListNode{4,nil}})
	//log.Println(s)
}
