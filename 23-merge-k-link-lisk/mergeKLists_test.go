package _3_merge_k_link_lisk

import (
	"log"
	"testing"
)

func TestMergeKLists(t *testing.T) {
	l1 := &ListNode{Val: 1, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5}}}
	l2 := &ListNode{Val: 1, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4}}}
	l3 := &ListNode{Val: 2, Next: &ListNode{Val: 6}}
	//var l4 *ListNode
	lists := []*ListNode{l1,l2,l3}
	s2 := MergeKLists(lists)
	log.Println(s2)
}
