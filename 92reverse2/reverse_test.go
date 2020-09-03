package _2reverse2

import (
	"log"
	"testing"
)

func TestReverseBetween(t *testing.T) {
	//s1 := &ListNode{1, &ListNode{2, nil}}
	s1 := &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5,nil}}}}}
	s2:=ReverseBetween(s1, 1, 2)
	log.Println(s2)
}
